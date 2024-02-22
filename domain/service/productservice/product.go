package productservice

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	uuid "github.com/satori/go.uuid"
)

type ProductService struct {
	cache contract.CacheManager
	data  contract.DataManager
}

func NewProductService(
	cache contract.CacheManager,
	data contract.DataManager,
) contract.ProductService {
	return ProductService{
		cache: cache,
		data:  data,
	}
}

func (s ProductService) Insert(product entity.Product) (string, error) {
	product.UUID = uuid.NewV4().String()

	category, err := s.data.Category().FindByUUID(product.Category.UUID)
	if err != nil {
		return "", err
	}

	product.Category = category
	product.CategoryID = category.ID

	_, err = s.data.Product().Insert(product)
	if err != nil {
		return "", err
	}

	s.cache.Invalidate("product:all")

	return product.UUID, nil
}

func (s ProductService) Find() ([]entity.Product, error) {
	cacheKey := "product:all"
	var products []entity.Product
	err := s.cache.GetStruct(cacheKey, &products)
	if err != nil {
		products, err = s.data.Product().Find()
		if err != nil {
			return nil, err
		}

		s.cache.SetStruct(cacheKey, &products)
	}

	return products, nil
}

func (s ProductService) GetProductByUUID(productUUID string) (entity.Product, error) {
	product, err := s.data.Product().FindByUUID(productUUID)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}
