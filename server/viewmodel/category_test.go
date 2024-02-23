package viewmodel

import (
	"fmt"
	"testing"

	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestCategoryInsertRequest_Validate(t *testing.T) {
	testCases := []struct {
		name           string
		request        CategoryInsertRequest
		expectedErrors []error
	}{
		{
			name:           "ValidRequest",
			request:        CategoryInsertRequest{Name: "Valid Category"},
			expectedErrors: nil,
		},
		{
			name:           "MissingName",
			request:        CategoryInsertRequest{Name: ""},
			expectedErrors: []error{fmt.Errorf("invalid-parameter-name")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := tc.request.Validate()
			assert.Equal(t, tc.expectedErrors, errors)
		})
	}
}

func TestCategoryInsertRequest_Parse(t *testing.T) {
	categoryName := "Test Category"
	expected := entity.Category{Name: categoryName}

	request := CategoryInsertRequest{Name: categoryName}
	result := request.Parse()

	assert.Equal(t, expected, result)
}

func TestParseCategoryInsertResponse(t *testing.T) {
	uuid := "test-uuid"
	expected := CategoryInsertResponse{UUID: uuid}

	result := ParseCategoryInsertResponse(uuid)

	assert.Equal(t, expected, result)
}
