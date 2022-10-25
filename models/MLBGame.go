package models

import (
	"time"

	"gorm.io/gorm"
)

type MLBGame struct {
	gorm.Model
	GameID      int       `json:"id" gorm:"uniqueIndex"`
	Link        string    `json:"link"`
	HomeID      uint      `json:"-"`
	Home        *MLBTeam  `json:"home" gorm:"-;"`
	VisitorID   uint      `json:"-"`
	Visitor     *MLBTeam  `json:"visitor" gorm:"-"`
	Description string    `json:"description"`
	Gametime    time.Time `json:"gametime"`
}

type MLBTeam struct {
	gorm.Model
	TeamID int    `json:"id" gorm:"uniqueIndex"`
	Name   string `json:"name"`
	Link   string `json:"link"`
}
