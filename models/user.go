package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

// User model.
// @todo: force unique Name.
type User struct {
	gorm.Model
	Name         string    `json:"username,omitempty"`
	Zips         []UserZip `json:"zipcodes"`
	Password     string    `json:"-" gorm:"-"`
	PasswordHash string    `json:"-"`
}

type UserZip struct {
	gorm.Model
	Zip    string
	UserID uint
}

// MarshalJSON converts struct to string value.
func (uz UserZip) MarshalJSON() ([]byte, error) {
	return json.Marshal(uz.Zip)
}

// GetUserName returns the name field.
func (u *User) GetUserName() string {
	return u.Name
}

// LoadByName uses name as key to find user in DB.
func (u *User) LoadByName(db Store) error {
	tx := db.First(&u, "name=?", u.Name)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Create generates a new user in storage.
func (u *User) Create(db Store) error {
	if u.Name == "" || u.Password == "" {
		return errors.New("username and password are required")
	}

	ph, err := u.hashPassword()
	if err != nil {
		return err
	}
	u.PasswordHash = ph

	// Add to storage, if entry does not exist.
	tx := db.Create(&u)

	if tx.Error != nil {
		return fmt.Errorf("failed to add user")
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
func (u *User) CheckPasswordHash(db Store, p string) bool {
	if u.ID == 0 {
		err := u.LoadByName(db)
		if err != nil {
			log.Print(err)
			return false
		}
	}
	tx := db.First(&u, "id=?", u.ID)
	if tx.Error != nil {
		// TODO: what happens here?
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(p))
	return err == nil
}

// GetZips looks up the user's zip codes in storage.
func (u *User) GetZips(db Store) error {
	if u.ID == 0 {
		err := u.LoadByName(db)
		if err != nil {
			log.Print(err)
			return errors.New("nil user")
		}
	}
	var uz []UserZip
	tx := db.Find(&uz, "user_id=?", u.ID)
	if tx.Error != nil {
		return errors.New("error user zips")
	}
	u.Zips = uz
	return nil
}

// HasZip checks if zip is stored in DB.
func (u *User) HasZip(db Store, zip string) (bool, error) {
	if u.ID == 0 {
		err := u.LoadByName(db)
		if err != nil {
			log.Print(err)
			return false, errors.New("nil user")
		}
	}
	var z UserZip
	tx := db.First(&z, "user_id=? AND zip=?", u.ID, zip)
	// No errors means zip is present in DB.
	if tx.Error == nil {
		return true, nil
	}
	// @todo: properly handle this error in this function?
	return false, tx.Error
}

// AddZip adds an entry to user's zip codes in storage.
func (u *User) AddZip(db Store, newZip string) error {
	if u.ID == 0 {
		err := u.LoadByName(db)
		if err != nil {
			log.Print(err)
			return errors.New("nil user")
		}
	}
	ok, _ := u.HasZip(db, newZip)
	if ok {
		return errors.New("user already has zip")
	}
	z := UserZip{Zip: newZip, UserID: u.ID}
	tx := db.Create(&z)
	if tx.Error != nil {
		return tx.Error
	}
	// Reload user to properly set zips.
	db.Preload("Zips").First(&u, u.ID)
	return nil
}

// DeleteZip removes a user's zip code from storage.
func (u *User) DeleteZip(db Store, zip string) error {
	if u.ID == 0 {
		err := u.LoadByName(db)
		if err != nil {
			log.Print(err)
			return errors.New("nil user")
		}
	}
	z := UserZip{}
	tx := db.First(&z, "user_id=? AND zip=?", u.ID, zip)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Delete(&z)
	if tx.Error != nil {
		return tx.Error
	}
	// Reload user to properly set zips.
	db.Preload("Zips").First(&u, u.ID)
	return nil
}

// ClearZips removes all user zip codes from storage.
func (u *User) ClearZips(db Store) error {
	if u.ID == 0 {
		err := u.LoadByName(db)
		if err != nil {
			log.Print(err)
			return errors.New("nil user")
		}
	}
	var zs []UserZip
	tx := db.Delete(&zs, "user_id=?", u.ID)
	if tx.Error != nil {
		return tx.Error
	}

	// Reload user to properly set zips.
	tx = db.Preload("Zips").First(&u, u.ID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
