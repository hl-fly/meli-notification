package entity

import "gorm.io/gorm"

type UserProductHistory struct {
	gorm.Model
	ProductID uint
	UserID    uint
	User      User
	Product   Product
}

type UsersProductHistories []UserProductHistory

func (e UsersProductHistories) GetUsers() []User {
	users := make([]User, 0)
	for _, history := range e {
		users = append(users, history.User)
	}

	return users
}
