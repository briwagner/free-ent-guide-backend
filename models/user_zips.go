package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"log"
	"slices"
)

// TODO we don't really need this from the outside API.
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
		log.Printf("doing user set zips with %s and %s", data, u.Email)
		err = q.UserSetZips(ctx, modelstore.UserSetZipsParams{
			JSONEXTRACT: string(data),
			Email:       u.Email,
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

	u.Data.Zips = make([]int64, 0)
	return nil
}
