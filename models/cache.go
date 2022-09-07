package models

import (
	"time"

	"gorm.io/gorm"
)

type Cache struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Value   string `gorm:"size:65000"`
	Expires time.Time
}
