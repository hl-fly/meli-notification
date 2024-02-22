package data

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"gorm.io/gorm"
)

type userProductHistoryRepo struct {
	db *gorm.DB
}

func newUserProductHistory(db *gorm.DB) contract.UserProductHistoryRepo {
	return userProductHistoryRepo{
		db: db,
	}
}

func (r userProductHistoryRepo) Insert(userProductHistory entity.UserProductHistory) (uint, error) {
	result := r.db.Create(&userProductHistory)
	if result.Error != nil {
		return 0, result.Error
	}

	return userProductHistory.ID, nil
}

func (r userProductHistoryRepo) FindByProductID(productID uint) ([]entity.UserProductHistory, error) {
	var userProductHistory []entity.UserProductHistory

	result := r.db.Joins("User").Where("product_id = ?", productID).Select("DISTINCT ON(user_id) *").Order("user_id").Find(&userProductHistory)
	if result.Error != nil {
		return []entity.UserProductHistory{}, result.Error
	}

	return userProductHistory, nil
}
