package contract

import (
	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/hector-leite/meli-notification/infra/context"
)

type AuthTokenService interface {
	NewTokenWrapper(user entity.User, tokenType string) (entity.TokenWrapper, error)
	FindUnexpiredByUUIDAndType(authTokenUUID string, authTokenType string) (entity.AuthToken, error)
	DeleteByID(authTokenID uint) error
}

type UserService interface {
	GetFromContext(context context.Context) (entity.User, error)
	Find() ([]entity.User, error)
	SignIn(password string, user entity.User) (entity.User, error)
	SignUp(password string, user entity.User) (err error)
	RefreshToken(hashAuthToken string) (entity.User, error)
}

type CategoryService interface {
	Insert(category entity.Category) (string, error)
}

type ProductService interface {
	Find() ([]entity.Product, error)
	Insert(product entity.Product) (string, error)
	GetProductByUUID(productUUID string) (entity.Product, error)
}

type UserProductHistoryService interface {
	Insert(userProductHistory entity.UserProductHistory) (err error)
	GetUsersByProductID(productID uint) (entity.UsersProductHistories, error)
}

type NotificationService interface {
	Insert(notification entity.Notification) (err error)
}

type UserNotificationService interface {
	FindByUser(user entity.User) ([]entity.UserNotification, error)
	Insert(userID, notificationID uint) (err error)
}
