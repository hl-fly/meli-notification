package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	UUID string `gorm:"not null"`
	Name string `gorm:"not null; uniqueIndex:uidx_name,sort:asc"`
}
