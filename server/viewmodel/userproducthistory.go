package viewmodel

import (
	"fmt"

	"github.com/hector-leite/meli-notification/domain/entity"
)

type UserProductHistoryRequest struct {
	ProductUUID string `json:"product_uuid"`
}

func (vm UserProductHistoryRequest) Validate() []error {
	genericErrorFormat := "invalid-parameter-%s"
	var aggErrors []error

	if vm.ProductUUID == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "product_uuid"))
	}

	if len(aggErrors) > 0 {
		return aggErrors
	}

	return nil
}

func (vm UserProductHistoryRequest) Parse(user entity.User) entity.UserProductHistory {
	product := entity.Product{
		UUID: vm.ProductUUID,
	}

	return entity.UserProductHistory{
		Product: product,
		UserID:  user.ID,
		User:    user,
	}
}
