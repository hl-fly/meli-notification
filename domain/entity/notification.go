package entity

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UUID           string    `gorm:"not null"`
	ExpirationDate time.Time `gorm:"not null"`
	ProductID      uint
	Message        string
	Link           string
	Target         bool    `gorm:"not null; default:false"`
	Product        Product `gorm:"-"`
}
