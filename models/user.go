package models

import (
	"errors"
	"fmt"

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

func (u *User) getDbKey() string {
	return fmt.Sprintf("user:%s", u.Name)
}

// GetUserName will return the name field.
func (u *User) GetUserName() string {
	return u.Name
}

// Create generates a new user in storage.
func (u *User) Create(db *redis.Client) error {
	if u.Name == "" || u.Password == "" {
		return errors.New("username and password are required")
	}

	ph, err := u.HashPassword()
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

	return nil
}

// HashPassword returns the hashed password.
func (u *User) HashPassword() (string, error) {
	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("password hash error")
	}

	return string(ph), nil
}

// CheckPasswordHash compares pass to stored hash.
func (u *User) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
