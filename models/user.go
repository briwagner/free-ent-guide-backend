package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

// User model.
type User struct {
	Name         string   `json:"username,omitempty"`
	Zips         []string `json:"zipcodes"`
	Password     string   `json:"-"`
	PasswordHash string   `json:"-"`
}

// getDbKey formats the Redis key for a user.
func (u *User) getDbKey() string {
	return fmt.Sprintf("user:%s", u.Name)
}

// GetUserName returns the name field.
func (u *User) GetUserName() string {
	return u.Name
}

// Create generates a new user in storage.
func (u *User) Create(db *redis.Client) error {
	if u.Name == "" || u.Password == "" {
		return errors.New("username and password are required")
	}

	ph, err := u.hashPassword()
	if err != nil {
		return err
	}
	u.PasswordHash = ph

	// Add to storage, if entry does not exist.
	ok, err := db.SetNX(u.getDbKey(), u.PasswordHash, 0).Result()
	if err != nil {
		return fmt.Errorf("failed to add user")
	}
	if !ok {
		return fmt.Errorf("user exists: %s", u.Name)
	}

	log.Printf("New user created %s", u.Name)

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
func (u *User) CheckPasswordHash(db *redis.Client, p string) bool {
	h, err := db.Get(fmt.Sprintf("user:%s", u.Name)).Result()
	if err != nil {
		// TODO: what happens here?
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}

// GetZips looks up the user's zip codes in storage.
func (u *User) GetZips(db *redis.Client) error {
	records, err := db.LRange(fmt.Sprintf("user:%s:zip", u.Name), 0, -1).Result()
	if err != nil {
		return errors.New("error uwer zips")
	}
	u.Zips = records
	return nil
}

// AddZip adds an entry to user's zip codes in storage.
func (u *User) AddZip(db *redis.Client, newZip string) error {
	res := db.RPush(fmt.Sprintf("user:%s:zip", u.Name), newZip)
	if res.Err() != nil {
		return errors.New("cannot add zip")
	}
	return nil
}

// DeleteZip removes a user's zip code from storage.
func (u *User) DeleteZip(db *redis.Client, zip string) error {
	res, err := db.LRem(fmt.Sprintf("user:%s:zip", u.Name), 0, zip).Result()
	if err != nil {
		return errors.New("cannot delete zip")
	}
	// TODO: restrict entering repeated values. Then change this to == 1.
	// How to pass this to frontend?
	// Redis returns # of items removed.
	if res >= 1 {
		log.Println("ok removing zip")
	} else {
		log.Println("zip not found")
	}
	return nil
}

// ClearZips removes all user zip codes from storage.
func (u *User) ClearZips(db *redis.Client) error {
	res := db.Del(fmt.Sprintf("user:%s:zip", u.Name))
	if res.Err() != nil {
		return errors.New("cannot clear zips")
	}
	return nil
}
