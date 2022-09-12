package models

import (
	"time"

	"gorm.io/gorm"
)

type Cache struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Value   string `gorm:"type:mediumblob"`
	Expires time.Time
}
