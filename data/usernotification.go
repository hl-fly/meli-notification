package data

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"gorm.io/gorm"
)

type userNotificationRepo struct {
	db *gorm.DB
}

func newUserNotificationRepo(db *gorm.DB) contract.UserNotificationRepo {
	return userNotificationRepo{
		db: db,
	}
}

func (r userNotificationRepo) Insert(userNotification entity.UserNotification) (uint, error) {
	result := r.db.Create(&userNotification)
	if result.Error != nil {
		return 0, result.Error
	}

	return userNotification.ID, nil
}

func (r userNotificationRepo) FindByUserID(userID uint) ([]entity.UserNotification, error) {
	var userNotifications []entity.UserNotification

	result := r.db.Joins("Notification").Where("user_id = ?", userID).Find(&userNotifications)
	if result.Error != nil {
		return []entity.UserNotification{}, result.Error
	}

	return userNotifications, nil
}
