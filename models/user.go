package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Zips  []int64 `json:"zips"`
	Shows []int64 `json:"shows"`
}

// NewUserData creates a new struct with field slices prepared.
// Important for creating users with proper JSON keys to use with database operations.
func NewUserData() UserData {
	return UserData{
		Zips:  make([]int64, 0),
		Shows: make([]int64, 0),
	}
}

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
	if len(u.Data.Shows) != 0 {
		// Frontend expects this prop key.
		vals["shows"] = u.Data.Shows
	}

	vals["email"] = u.Email

	return json.Marshal(vals)
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
	// We need to init with `key => [empty slices]` to allow JSON database operations later.
	userdata, err := json.Marshal(NewUserData())
	if err != nil {
		return fmt.Errorf("error marshaling empty userdata %s", err)
	}

	id, err := q.UserCreate(context.Background(), modelstore.UserCreateParams{
		Email:        u.Email,
		PasswordHash: ph,
		Data:         userdata,
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

	userData := NewUserData()
	err = json.Unmarshal(row.Data, &userData)
	if err != nil {
		return fmt.Errorf("error parsing user data: %s", err)
	}
	u.Data = userData

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

// LoadData populates zips, shows, etc. on User.
func (u *User) LoadData(q *modelstore.Queries) error {
	if u.Email == "" {
		return errors.New("user email not set")
	}
	err := u.FindByEmail(q)
	if err != nil {
		return fmt.Errorf("error user find by email: %s", err)
	}

	return nil
}
