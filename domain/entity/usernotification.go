package entity

import (
	"gorm.io/gorm"
)

type UserNotification struct {
	gorm.Model
	UUID           string
	UserID         uint
	User           User
	NotificationID uint
	Notification   Notification
	IsRead         bool `gorm:"not null; default:false"`
}
