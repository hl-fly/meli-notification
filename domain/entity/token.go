package entity

import (
	"time"

	"gorm.io/gorm"
)

type AuthToken struct {
	gorm.Model
	UUID           string
	HashToken      []byte
	Type           string
	ExpirationDate time.Time
	UserID         *uint
}

type TokenWrapper struct {
	HashAuthToken string
	Type          string
}
