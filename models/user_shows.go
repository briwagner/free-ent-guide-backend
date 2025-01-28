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

// GetShows looks up the user's stored show IDs.
func (u *User) GetShows(q *modelstore.Queries) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	res, err := q.UserExtractShows(context.Background(), u.Email)
	if err != nil {
		log.Printf("error extracting shows %s", err)
		return err
	}

	data, ok := res.([]byte)
	if ok {
		var shows []int64
		err := json.Unmarshal(data, &shows)
		if err != nil {
			log.Println(err)
			return err
		}
		u.Data = UserData{Shows: shows}
	}

	return nil
}

// AddShow adds an entry to user's stored Shows.
func (u *User) AddShow(q *modelstore.Queries, newShow int64) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	err := q.UserAppendShow(context.Background(), modelstore.UserAppendShowParams{
		JSONARRAYAPPEND: newShow,
		Email:           u.Email,
	})
	if err != nil {
		return fmt.Errorf("error adding show %w", err)
	}

	u.Data.Shows = append(u.Data.Shows, newShow)
	return err
}

// DeleteShow removes a stored show.
func (u *User) DeleteShow(q *modelstore.Queries, show int64) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	if len(u.Data.Shows) == 0 {
		err := u.GetShows(q)
		if err != nil {
			return err
		}

		if len(u.Data.Shows) == 0 {
			return errors.New("no shows stored")
		}
	}

	if !slices.Contains(u.Data.Shows, show) {
		return fmt.Errorf("show not found: %d", show)
	}

	ctx := context.Background()
	var newShows []int64
	for _, s := range u.Data.Shows {
		if s != show {
			newShows = append(newShows, s)
		}
	}

	data, err := json.Marshal(newShows)
	if err != nil {
		return fmt.Errorf("error marshalling user shows: %w", err)
	}
	if len(newShows) == 0 {
		err = q.UserClearShows(ctx, u.Email)
	} else {
		err = q.UserSetShows(ctx, modelstore.UserSetShowsParams{
			JSONEXTRACT: string(data),
			Email:       u.Email,
		})
	}

	if err != nil {
		return err
	}

	u.Data.Shows = newShows
	return nil
}

// ClearShows removes all stored user shows.
func (u *User) ClearShows(q *modelstore.Queries) error {
	if u.Email == "" {
		return errors.New("invalid user email")
	}

	err := q.UserClearShows(context.Background(), u.Email)
	if err != nil {
		return err
	}

	u.Data.Shows = make([]int64, 0)
	return nil
}
