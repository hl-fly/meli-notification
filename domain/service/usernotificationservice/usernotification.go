package usernotificationservice

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type UserNotificationService struct {
	data contract.DataManager
}

func NewUserNotificationService(
	data contract.DataManager,
) contract.UserNotificationService {
	return UserNotificationService{
		data: data,
	}
}

func (s UserNotificationService) Insert(userID, notificationID uint) (err error) {
	userNottification := entity.UserNotification{
		UUID:           uuid.NewV4().String(),
		UserID:         userID,
		NotificationID: notificationID,
	}

	_, err = s.data.UserNotification().Insert(userNottification)
	if err != nil {
		return err
	}

	return nil
}

func (s UserNotificationService) FindByUser(user entity.User) ([]entity.UserNotification, error) {
	userNotifications, err := s.data.UserNotification().FindByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	return userNotifications, nil
}
