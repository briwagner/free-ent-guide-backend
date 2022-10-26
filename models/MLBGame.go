package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MLBGame struct {
	gorm.Model
	GameID      int       `json:"id" gorm:"uniqueIndex"`
	Link        string    `json:"link"`
	HomeID      uint      `json:"-"`
	Home        MLBTeam   `json:"home"`
	VisitorID   uint      `json:"-"`
	Visitor     MLBTeam   `json:"visitor"`
	Description string    `json:"description"`
	Gametime    time.Time `json:"gametime"`
}

type MLBTeam struct {
	gorm.Model
	TeamID int    `json:"id" gorm:"uniqueIndex"`
	Name   string `json:"name"`
	Link   string `json:"link"`
}

type MLBGames []MLBGame

// LoadByDate fetches all games for the given date.
func (mgs *MLBGames) LoadByDate(d string, db Store) error {
	d = fmt.Sprintf("%s%%", d)
	tx := db.Preload("Home").Preload("Visitor").Where("gametime LIKE ?", d).Find(mgs)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindByID searches for a game by Game ID.
func (g *MLBGame) FindByID(id string, db Store) error {
	tx := db.Preload("Home").Preload("Visitor").Where("game_id = ?", id).Find(g)
	if tx.Error != nil {
		return tx.Error
	}
	if g.ID == 0 {
		return errors.New("MLB game not found")
	}
	return nil
}
