package data

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"gorm.io/gorm"
)

type notificationRepo struct {
	db *gorm.DB
}

func newNotificationRepo(db *gorm.DB) contract.NotificationRepo {
	return notificationRepo{
		db: db,
	}
}

func (r notificationRepo) Insert(notification entity.Notification) (uint, error) {
	result := r.db.Create(&notification)
	if result.Error != nil {
		return 0, result.Error
	}

	return notification.ID, nil
}
