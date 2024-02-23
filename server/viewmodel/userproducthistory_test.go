package viewmodel

import (
	"fmt"
	"testing"

	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserProductHistoryRequest_Validate(t *testing.T) {
	testCases := []struct {
		name           string
		request        UserProductHistoryRequest
		expectedErrors []error
	}{
		{
			name:           "ValidRequest",
			request:        UserProductHistoryRequest{ProductUUID: "valid-uuid"},
			expectedErrors: nil,
		},
		{
			name:           "InvalidRequest",
			request:        UserProductHistoryRequest{ProductUUID: ""},
			expectedErrors: []error{fmt.Errorf("invalid-parameter-product_uuid")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := tc.request.Validate()
			assert.Equal(t, tc.expectedErrors, errors)
		})
	}
}

func TestUserProductHistoryRequest_Parse(t *testing.T) {
	user := entity.User{Model: gorm.Model{ID: 123}, UUID: "test-uuid", CPF: "12345678901", Name: "John Doe", Email: "john@example.com", AllowNotification: true, Type: "user"}
	productUUID := "test-uuid"
	expected := entity.UserProductHistory{
		Product: entity.Product{UUID: productUUID},
		UserID:  user.ID,
		User:    user,
	}

	request := UserProductHistoryRequest{ProductUUID: productUUID}
	result := request.Parse(user)

	assert.Equal(t, expected, result)
}
