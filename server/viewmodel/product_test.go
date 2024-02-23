package viewmodel

import (
	"fmt"
	"testing"

	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestProductInsertRequest_Validate(t *testing.T) {
	testCases := []struct {
		name           string
		request        ProductInsertRequest
		expectedErrors []error
	}{
		{
			name:           "ValidRequest",
			request:        ProductInsertRequest{Name: "Valid Product", CategoryUUID: "valid-uuid"},
			expectedErrors: nil,
		},
		{
			name:           "MissingName",
			request:        ProductInsertRequest{Name: "", CategoryUUID: "valid-uuid"},
			expectedErrors: []error{fmt.Errorf("invalid-parameter-name")},
		},
		{
			name:           "MissingCategoryUUID",
			request:        ProductInsertRequest{Name: "Valid Product", CategoryUUID: ""},
			expectedErrors: []error{fmt.Errorf("invalid-parameter-category_uuid")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := tc.request.Validate()
			assert.Equal(t, tc.expectedErrors, errors)
		})
	}
}

func TestProductInsertRequest_Parse(t *testing.T) {
	categoryUUID := "test-category-uuid"
	productName := "Test Product"
	expected := entity.Product{Name: productName, Category: entity.Category{UUID: categoryUUID}}

	request := ProductInsertRequest{Name: productName, CategoryUUID: categoryUUID}
	result := request.Parse()

	assert.Equal(t, expected, result)
}

func TestParseProductInsertResponse(t *testing.T) {
	uuid := "test-uuid"
	expected := ProductInsertResponse{UUID: uuid}

	result := ParseProductInsertResponse(uuid)

	assert.Equal(t, expected, result)
}

func TestParseGetProductsResponse(t *testing.T) {
	products := []entity.Product{
		{UUID: "uuid1", Name: "Product 1"},
		{UUID: "uuid2", Name: "Product 2"},
	}

	expected := GetProductsResponse{
		{UUID: "uuid1", Name: "Product 1"},
		{UUID: "uuid2", Name: "Product 2"},
	}

	result := ParseGetProductsResponse(products)

	assert.Equal(t, expected, result)
}
