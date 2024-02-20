package models

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
)

// User model.
type User struct {
	ID           int    `json:"id"`
	Email        string `json:"username,omitempty"`
	Data         UserData
	ResetCode    string    `json:"-"`
	Password     string    `json:"-"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"updated_at"`
	Status       bool      `json:"-"`
}

func (u User) MarshalJSON() ([]byte, error) {
	vals := make(map[string]interface{})
	vals["ID"] = u.ID

	if len(u.Data.Zips) != 0 {
		// Frontend expects this prop key.
		vals["zipcodes"] = u.Data.Zips
	}

	vals["email"] = u.Email

	return json.Marshal(vals)
}

type UserData struct {
	Zips []int64 `json:"zips"`
}

func (ud *UserData) UnmarshalJSON(data []byte) error {
	var vals map[string]interface{}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	err := dec.Decode(&vals)
	if err != nil {
		return err
	}

	rawZips, ok := vals["zips"]
	if !ok {
		fmt.Println("no zips")
		return nil
	}
	zips := rawZips.([]interface{})
	if len(zips) != 0 {
		userZips := make([]int64, len(zips))
		for i, z := range zips {
			jn, err := z.(json.Number).Int64()
			if err != nil {
				log.Printf("error parsing zips %s", err)
				continue
			}
			userZips[i] = jn
		}
		ud.Zips = userZips
	}
	return nil
}

// Create generates a new user in storage.
func (u *User) Create(q *modelstore.Queries) error {
	if u.Email == "" || u.Password == "" {
		return errors.New("username and password are required")
	}

	ph, err := u.hashPassword()
	if err != nil {
		return err
	}
	u.PasswordHash = ph

	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	data := `{"zips":[]}` // if we don't set this now, we can't append cuz it doesn't exist.

	id, err := q.UserCreate(context.Background(), modelstore.UserCreateParams{
		Email:        u.Email,
		PasswordHash: ph,
		Data:         []byte(data),
		CreatedAt:    sql.NullTime{Time: u.CreatedAt, Valid: true},
	})

	if err != nil {
		return fmt.Errorf("failed to add user: %w", err)
	}
	u.ID = int(id)
	return nil
}

func (u *User) FindByEmail(q *modelstore.Queries) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	row, err := q.UserFindByEmail(context.Background(), u.Email)
	if err != nil {
		return fmt.Errorf("error finding user by email: %w", err)
	}

	u.ID = int(row.ID)
	u.CreatedAt = row.CreatedAt.Time
	u.Email = row.Email
	u.Status = row.Status.Bool

	if row.Zips != nil {
		data, ok := row.Zips.([]byte)
		if ok {
			var zips []int64
			err := json.Unmarshal(data, &zips)
			if err != nil {
				log.Println(err)
				return err
			}
			u.Data = UserData{Zips: zips}
		}
	}

	return nil
}

// hashPassword returns the hashed password.
func (u *User) hashPassword() (string, error) {
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("password hash error")
	}

	return string(ph), nil
}

// CheckPasswordHash compares pass to stored hash.
func (u *User) CheckPasswordHash(q *modelstore.Queries, p string) bool {
	if u.Email == "" {
		log.Printf("user email not set")
		return false
	}
	if u.ID == 0 {
		err := u.FindByEmail(q)
		if err != nil {
			log.Printf("error user find by email: %s", err)
			return false
		}
	}

	if !u.Status {
		log.Printf("user is not active: %s", u.Email)
		return false
	}

	rows, err := q.UserCheckPassword(context.Background(), u.Email)
	if err != nil {
		log.Printf("error checking pw hash %s", err)
		return false
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(rows.PasswordHash), []byte(p))
	return err2 == nil
}

// GetZips looks up the user's zip codes in storage.
func (u *User) GetZips(q *modelstore.Queries) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	res, err := q.UserExtractZips(context.Background(), u.Email)
	if err != nil {
		log.Printf("error extracting zips %s", err)
		return err
	}

	data, ok := res.([]byte)
	if ok {
		var zips []int64
		err := json.Unmarshal(data, &zips)
		if err != nil {
			log.Println(err)
			return err
		}
		u.Data = UserData{Zips: zips}
	}

	return nil
}

// AddZip adds an entry to user's zip codes in storage.
func (u *User) AddZip(q *modelstore.Queries, newZip int64) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	// TODO validate that it's an int with six digits.

	// TODO this does not check if zip exists.
	err := q.UserAppendZip(context.Background(), modelstore.UserAppendZipParams{
		JSONARRAYAPPEND: newZip,
		Email:           u.Email,
	})
	if err != nil {
		return fmt.Errorf("error adding zip %w", err)
	}

	u.Data.Zips = append(u.Data.Zips, newZip)
	return err
}

// DeleteZip removes a user's zip code from storage.
func (u *User) DeleteZip(q *modelstore.Queries, zip int64) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	if len(u.Data.Zips) == 0 {
		err := u.GetZips(q)
		if err != nil {
			return err
		}

		if len(u.Data.Zips) == 0 {
			return errors.New("no zips stored")
		}
	}

	if !slices.Contains(u.Data.Zips, zip) {
		return fmt.Errorf("zip not found: %d", zip)
	}

	ctx := context.Background()
	var newZips []int64
	for _, z := range u.Data.Zips {
		if z != zip {
			newZips = append(newZips, z)
		}
	}

	data, err := json.Marshal(newZips)
	if err != nil {
		return fmt.Errorf("error marshalling user zips: %w", err)
	}
	if len(newZips) == 0 {
		err = q.UserClearZips(ctx, u.Email)
	} else {
		err = q.UserSetZips(ctx, modelstore.UserSetZipsParams{
			Column1: data,
			Email:   u.Email,
		})
	}

	if err != nil {
		return err
	}

	u.Data.Zips = newZips
	return nil
}

// ClearZips removes all user zip codes from storage.
func (u *User) ClearZips(q *modelstore.Queries) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	err := q.UserClearZips(context.Background(), u.Email)
	if err != nil {
		return err
	}

	u.Data.Zips = []int64{}
	return nil
}
