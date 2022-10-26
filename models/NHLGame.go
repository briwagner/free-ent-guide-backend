package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type NHLGame struct {
	gorm.Model
	GameID      int       `json:"id" gorm:"uniqueIndex"`
	Link        string    `json:"link"`
	HomeID      uint      `json:"-"`
	Home        NHLTeam   `json:"home"`
	VisitorID   uint      `json:"-"`
	Visitor     NHLTeam   `json:"visitor"`
	Description string    `json:"description"`
	Gametime    time.Time `json:"gametime"`
}

type NHLTeam struct {
	gorm.Model
	TeamID int    `json:"id" gorm:"uniqueIndex"`
	Name   string `json:"name"`
	Link   string `json:"link"`
}

type NHLGames []NHLGame

// LoadByDate fetches all games for the given date.
func (ngs *NHLGames) LoadByDate(d string, db Store) error {
	d = fmt.Sprintf("%s%%", d)
	tx := db.Preload("Home").Preload("Visitor").Where("gametime LIKE ?", d).Find(ngs)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindByID searches for a game by Game ID.
func (g *NHLGame) FindByID(id string, db Store) error {
	tx := db.Preload("Home").Preload("Visitor").Where("game_id = ?", id).Find(g)
	if tx.Error != nil {
		return tx.Error
	}
	if g.ID == 0 {
		return errors.New("NHL game not found")
	}
	return nil
}
