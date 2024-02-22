package notificationservice

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type NotificationService struct {
	data                      contract.DataManager
	productService            contract.ProductService
	userProductHistoryService contract.UserProductHistoryService
	userNotificationService   contract.UserNotificationService
}

func NewNotificationService(
	data contract.DataManager,
	productService contract.ProductService,
	userProductHistoryService contract.UserProductHistoryService,
	userNotificationService contract.UserNotificationService,
) contract.NotificationService {
	return NotificationService{
		data:                      data,
		productService:            productService,
		userProductHistoryService: userProductHistoryService,
		userNotificationService:   userNotificationService,
	}
}

func (s NotificationService) Insert(notification entity.Notification) (err error) {
	var users []entity.User
	var product entity.Product

	if notification.Product.UUID != "" {
		product, err = s.productService.GetProductByUUID(notification.Product.UUID)
		if err != nil {
			return err
		}
		if notification.Target {
			usersHistory, err := s.userProductHistoryService.GetUsersByProductID(product.ID)
			if err != nil {
				return err
			}
			users = usersHistory.GetUsers()
		}

	}

	if len(users) == 0 {
		users, err = s.data.User().Find()
		if err != nil {
			return err
		}
	}

	notification.UUID = uuid.NewV4().String()
	notification.ProductID = product.ID
	notificationID, err := s.data.Notification().Insert(notification)
	if err != nil {
		return err
	}

	for _, user := range users {
		err = s.userNotificationService.Insert(user.ID, notificationID)
		if err != nil {
			return err
		}
	}

	return nil
}
