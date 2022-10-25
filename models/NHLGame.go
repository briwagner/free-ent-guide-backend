package models

import (
	"time"

	"gorm.io/gorm"
)

type NHLGame struct {
	gorm.Model
	GameID      int       `json:"id" gorm:"uniqueIndex"`
	Link        string    `json:"link"`
	HomeID      uint      `json:"-"`
	Home        *NHLTeam  `json:"home" gorm:"-;"`
	VisitorID   uint      `json:"-"`
	Visitor     *NHLTeam  `json:"visitor" gorm:"-"`
	Description string    `json:"description"`
	Gametime    time.Time `json:"gametime"`
}

type NHLTeam struct {
	gorm.Model
	TeamID int    `json:"id" gorm:"uniqueIndex"`
	Name   string `json:"name"`
	Link   string `json:"link"`
}
