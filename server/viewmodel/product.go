package viewmodel

import (
	"fmt"

	"github.com/hector-leite/meli-notification/domain/entity"
)

type ProductInsertRequest struct {
	Name         string `json:"name"`
	CategoryUUID string `json:"category_uuid"`
}

func (vm ProductInsertRequest) Validate() []error {
	genericErrorFormat := "invalid-parameter-%s"
	var aggErrors []error

	if vm.Name == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "name"))
	}

	if vm.CategoryUUID == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "category_uuid"))
	}

	if len(aggErrors) > 0 {
		return aggErrors
	}

	return nil
}

func (vm ProductInsertRequest) Parse() entity.Product {
	category := entity.Category{
		UUID: vm.CategoryUUID,
	}

	return entity.Product{
		Name:     vm.Name,
		Category: category,
	}
}

type ProductInsertResponse struct {
	UUID string `json:"uuid"`
}

func ParseProductInsertResponse(UUID string) ProductInsertResponse {
	return ProductInsertResponse{
		UUID: UUID,
	}
}

type GetProductsResponse []GetProductResponse

type GetProductResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func ParseGetProductsResponse(products []entity.Product) GetProductsResponse {
	productsResponse := make([]GetProductResponse, len(products))

	for i := 0; i < len(products); i++ {
		productsResponse[i].UUID = products[i].UUID
		productsResponse[i].Name = products[i].Name
	}
	return productsResponse
}
