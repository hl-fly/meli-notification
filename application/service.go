package application

import (
	"github.com/hector-leite/meli-notification/domain/contract"
)

type AppService struct {
	authTokenService          contract.AuthTokenService
	userService               contract.UserService
	categoryService           contract.CategoryService
	productService            contract.ProductService
	userProductHistoryService contract.UserProductHistoryService
	notificationService       contract.NotificationService
	userNotificationService   contract.UserNotificationService
}

func NewAppService(
	authTokenService contract.AuthTokenService,
	userService contract.UserService,
	categoryService contract.CategoryService,
	productService contract.ProductService,
	userProductHistoryService contract.UserProductHistoryService,
	notificationService contract.NotificationService,
	userNotificationService contract.UserNotificationService,
) *AppService {
	return &AppService{
		authTokenService:          authTokenService,
		userService:               userService,
		categoryService:           categoryService,
		productService:            productService,
		userProductHistoryService: userProductHistoryService,
		notificationService:       notificationService,
		userNotificationService:   userNotificationService,
	}
}

func (svc AppService) AuthToken() contract.AuthTokenService {
	return svc.authTokenService
}

func (svc AppService) User() contract.UserService {
	return svc.userService
}

func (svc AppService) Category() contract.CategoryService {
	return svc.categoryService
}

func (svc AppService) Product() contract.ProductService {
	return svc.productService
}

func (svc AppService) UserProductHistory() contract.UserProductHistoryService {
	return svc.userProductHistoryService
}

func (svc AppService) Notification() contract.NotificationService {
	return svc.notificationService
}

func (svc AppService) UserNotification() contract.UserNotificationService {
	return svc.userNotificationService
}
