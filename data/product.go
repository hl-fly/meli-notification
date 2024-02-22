package data

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func newProductRepo(db *gorm.DB) contract.ProductRepo {
	return productRepo{
		db: db,
	}
}

func (r productRepo) Insert(product entity.Product) (uint, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}

	return product.ID, nil
}

func (r productRepo) FindByUUID(UUID string) (entity.Product, error) {
	var product entity.Product

	result := r.db.Where("uuid = ?", UUID).Find(&product)
	if result.Error != nil {
		return entity.Product{}, result.Error
	}

	return product, nil
}

func (r productRepo) Find() ([]entity.Product, error) {
	var products []entity.Product

	result := r.db.Find(&products)
	if result.Error != nil {
		return []entity.Product{}, result.Error
	}

	return products, nil
}
