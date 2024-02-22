package viewmodel

import (
	"fmt"

	"github.com/hector-leite/meli-notification/domain/entity"
)

type CategoryInsertRequest struct {
	Name string `json:"name"`
}

func (vm CategoryInsertRequest) Validate() []error {
	genericErrorFormat := "invalid-parameter-%s"
	var aggErrors []error

	if vm.Name == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "name"))
	}

	if len(aggErrors) > 0 {
		return aggErrors
	}

	return nil
}

func (vm CategoryInsertRequest) Parse() entity.Category {
	return entity.Category{
		Name: vm.Name,
	}
}

type CategoryInsertResponse struct {
	UUID string `json:"uuid"`
}

func ParseCategoryInsertResponse(UUID string) CategoryInsertResponse {
	return CategoryInsertResponse{
		UUID: UUID,
	}
}
