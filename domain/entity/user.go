package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID              string `gorm:"not null"`
	CPF               string `gorm:"not null; uniqueIndex:uidx_cpf,sort:asc; size:11"`
	Name              string `gorm:"not null; uniqueIndex:uidx_name,sort:asc"`
	Email             string `gorm:"not null; uniqueIndex:uidx_email,sort:asc; size:50"`
	AllowNotification bool   `gorm:"not null; default:true"`
	UserNotification  []UserNotification
	HashPassword      []byte         `gorm:"not null"`
	Type              string         `gorm:"not null; default:user"`
	Tokens            []TokenWrapper `gorm:"-"`
}
