package viewmodel

import (
	"fmt"
	"testing"
	"time"

	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNotificationInsertRequest_Validate(t *testing.T) {
	testCases := []struct {
		name           string
		request        NotificationInsertRequest
		expectedErrors []error
	}{
		{
			name: "ValidRequest",
			request: NotificationInsertRequest{
				ProductUUID: "valid-uuid",
				Message:     "Valid Message",
				Link:        "http://example.com",
				ExpDate:     time.Now().Add(24 * time.Hour),
				Target:      true,
			},
			expectedErrors: nil,
		},
		{
			name: "MissingMessage",
			request: NotificationInsertRequest{
				ProductUUID: "valid-uuid",
				Message:     "",
				Link:        "http://example.com",
				ExpDate:     time.Now().Add(24 * time.Hour),
				Target:      true,
			},
			expectedErrors: []error{fmt.Errorf("invalid-parameter-message")},
		},
		{
			name: "ZeroExpDate",
			request: NotificationInsertRequest{
				ProductUUID: "valid-uuid",
				Message:     "Valid Message",
				Link:        "http://example.com",
				ExpDate:     time.Time{},
				Target:      true,
			},
			expectedErrors: []error{fmt.Errorf("invalid-parameter-exp_date")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := tc.request.Validate()
			assert.Equal(t, tc.expectedErrors, errors)
		})
	}
}

func TestNotificationInsertRequest_Parse(t *testing.T) {
	productUUID := "test-product-uuid"
	message := "Test Message"
	link := "http://example.com"
	expDate := time.Now().Add(24 * time.Hour)
	target := true

	expected := entity.Notification{
		Product: entity.Product{
			UUID: productUUID,
		},
		Message:        message,
		Link:           link,
		ExpirationDate: expDate,
		Target:         target,
	}

	request := NotificationInsertRequest{
		ProductUUID: productUUID,
		Message:     message,
		Link:        link,
		ExpDate:     expDate,
		Target:      target,
	}
	result := request.Parse()

	assert.Equal(t, expected, result)
}
