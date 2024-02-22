package data

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func newCategoryRepo(db *gorm.DB) contract.CategoryRepo {
	return categoryRepo{
		db: db,
	}
}

func (r categoryRepo) Insert(category entity.Category) (uint, error) {
	result := r.db.Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}

	return category.ID, nil
}

func (r categoryRepo) FindByUUID(UUID string) (entity.Category, error) {
	var category entity.Category

	result := r.db.Where("uuid = ?", UUID).Find(&category)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}

	return category, nil
}
