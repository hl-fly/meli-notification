package contract

import (
	"github.com/hector-leite/meli-notification/domain/entity"
)

type DataManager interface {
	User() UserRepo
	AuthToken() AuthTokenRepo
	Category() CategoryRepo
	Product() ProductRepo
	UserProductHistory() UserProductHistoryRepo
	Notification() NotificationRepo
	UserNotification() UserNotificationRepo
}

type AuthTokenRepo interface {
	FindUnexpiredByIDAndType(authTokenID uint, authTokenType string) (entity.AuthToken, error)
	Insert(authToken entity.AuthToken) (uint, error)
	FindUnexpiredByUUIDAndType(authTokenUUID string, authTokenType string) (entity.AuthToken, error)
	DeleteByID(authTokenID uint) error
}

type UserRepo interface {
	Find() ([]entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByUUID(UUID string) (entity.User, error)
	SignUp(user entity.User) (uint, error)
}

type CategoryRepo interface {
	Insert(category entity.Category) (uint, error)
	FindByUUID(UUID string) (entity.Category, error)
}

type ProductRepo interface {
	Find() ([]entity.Product, error)
	FindByUUID(UUID string) (entity.Product, error)
	Insert(category entity.Product) (uint, error)
}

type UserProductHistoryRepo interface {
	Insert(notification entity.UserProductHistory) (uint, error)
	FindByProductID(productID uint) ([]entity.UserProductHistory, error)
}

type NotificationRepo interface {
	Insert(notification entity.Notification) (uint, error)
}

type UserNotificationRepo interface {
	FindByUserID(userID uint) ([]entity.UserNotification, error)
	Insert(notification entity.UserNotification) (uint, error)
}
