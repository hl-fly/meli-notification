package data

import (
	"fmt"

	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/hector-leite/meli-notification/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Conn struct {
	db *gorm.DB
}

func Connect(dbConfig config.DBConfig) (contract.DataManager, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		dbConfig.Address, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn := new(Conn)
	conn.db = db

	conn.db.AutoMigrate(
		&entity.AuthToken{},
		&entity.User{},
		&entity.Category{},
		&entity.Product{},
		&entity.Notification{},
		&entity.UserNotification{},
		&entity.UserProductHistory{},
	)

	return conn, nil
}

func (c *Conn) User() contract.UserRepo {
	return newUserRepo(c.db)
}

func (c *Conn) AuthToken() contract.AuthTokenRepo {
	return newAuthTokenRepo(c.db)
}

func (c *Conn) Category() contract.CategoryRepo {
	return newCategoryRepo(c.db)
}

func (c *Conn) Product() contract.ProductRepo {
	return newProductRepo(c.db)
}

func (c *Conn) UserProductHistory() contract.UserProductHistoryRepo {
	return newUserProductHistory(c.db)
}

func (c *Conn) Notification() contract.NotificationRepo {
	return newNotificationRepo(c.db)
}

func (c *Conn) UserNotification() contract.UserNotificationRepo {
	return newUserNotificationRepo(c.db)
}
