package userproducthistoryservice

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
)

type UserProductHistoryService struct {
	data contract.DataManager
}

func NewUserProductHistoryService(
	data contract.DataManager,
) contract.UserProductHistoryService {
	return UserProductHistoryService{
		data: data,
	}
}

func (s UserProductHistoryService) Insert(userProductHistory entity.UserProductHistory) (err error) {
	product, err := s.data.Product().FindByUUID(userProductHistory.Product.UUID)
	if err != nil {
		return err
	}

	userProductHistory.Product = product
	userProductHistory.ProductID = product.ID

	_, err = s.data.UserProductHistory().Insert(userProductHistory)
	if err != nil {
		return err
	}

	return nil
}

func (s UserProductHistoryService) GetUsersByProductID(productID uint) (entity.UsersProductHistories, error) {
	userProductHistory, err := s.data.UserProductHistory().FindByProductID(productID)
	if err != nil {
		return []entity.UserProductHistory{}, err
	}

	return userProductHistory, nil
}
