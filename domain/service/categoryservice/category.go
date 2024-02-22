package categoryservice

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type CategoryService struct {
	data contract.DataManager
}

func NewCategoryService(
	data contract.DataManager,
) contract.CategoryService {
	return CategoryService{
		data: data,
	}
}

func (s CategoryService) Insert(category entity.Category) (string, error) {
	category.UUID = uuid.NewV4().String()

	_, err := s.data.Category().Insert(category)
	if err != nil {
		return "", err
	}

	return category.UUID, nil
}
