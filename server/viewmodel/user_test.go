package viewmodel

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserSignupValidate(t *testing.T) {
	genericErrorFormat := "invalid-parameter-%s"
	for _, tc := range []struct {
		name              string
		signUpUserRequest SignUpUserRequest
		expectedErrors    []error
	}{
		{
			name: "ShouldSucessfullyValidateAll",
			signUpUserRequest: SignUpUserRequest{
				Name:            "Alien",
				Password:        "asdasd123",
				ConfirmPassword: "asdasd123",
				CPF:             "01234567890",
				Email:           "alien@gmail.com",
			},
		},
		{
			name: "ShouldReturnErrorWhenFieldsEmpty",
			signUpUserRequest: SignUpUserRequest{
				Name:            "",
				Password:        "",
				ConfirmPassword: "",
				CPF:             "",
				Email:           "",
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "name"),
				fmt.Errorf(genericErrorFormat, "email"),
				fmt.Errorf(genericErrorFormat, "cpf"),
				fmt.Errorf(genericErrorFormat, "password"),
				fmt.Errorf(genericErrorFormat, "confirm-password"),
			},
		},
		{
			name: "ShouldReturnErrorWhenPasswordsNotEqual",
			signUpUserRequest: SignUpUserRequest{
				Name:            "Alien",
				Password:        "asdasd123",
				ConfirmPassword: "asdasd124",
				CPF:             "01234567890",
				Email:           "alien@gmail.com",
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "password-confirm-password"),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErrors, tc.signUpUserRequest.Validate())
		})
	}
}

func TestUserParse(t *testing.T) {
	for _, tc := range []struct {
		name              string
		signUpUserRequest SignUpUserRequest
		expected          entity.User
	}{
		{
			name: "ShouldSucessfullyParseAll",
			signUpUserRequest: SignUpUserRequest{
				Name:  "Alien",
				CPF:   "01234567890",
				Email: "alien@mars.mar",
			},
			expected: entity.User{
				Name:  "Alien",
				CPF:   "01234567890",
				Email: "alien@mars.mar",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.signUpUserRequest.Parse())
		})
	}
}

func TestValidateUserSignin(t *testing.T) {
	genericErrorFormat := "invalid-parameter-%s"
	for _, tc := range []struct {
		name              string
		signInUserRequest SignInUserRequest
		expectedErrors    []error
	}{
		{
			name: "ShouldSucessfullyValidateAll",
			signInUserRequest: SignInUserRequest{
				Email:    "Alien",
				Password: "asdasd123",
			},
		},
		{
			name: "ShouldReturnErrorWhenFieldsEmpty",
			signInUserRequest: SignInUserRequest{
				Email:    "",
				Password: "",
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "email"),
				fmt.Errorf(genericErrorFormat, "password"),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErrors, tc.signInUserRequest.Validate())
		})
	}
}

func TestParseSignInUserRequest(t *testing.T) {
	for _, tc := range []struct {
		name              string
		signInUserRequest SignInUserRequest
		expected          entity.User
	}{
		{
			name: "ShouldSucessfullyParseAll",
			signInUserRequest: SignInUserRequest{
				Email: "alien@mars.mar",
			},
			expected: entity.User{
				Email: "alien@mars.mar",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.signInUserRequest.Parse())
		})
	}
}

func TestParseSignInUserRequestResponse(t *testing.T) {
	for _, tc := range []struct {
		name              string
		signInUserRequest SignInUserRequest
		user              entity.User
		expected          SignInUserRequestResponse
	}{
		{
			name:              "ShouldSucessfullyParseAll",
			signInUserRequest: SignInUserRequest{},
			user: entity.User{
				UUID: "123",
				Name: "Name",
				Tokens: []entity.TokenWrapper{
					{
						Type:          constants.TokenTypeAccess,
						HashAuthToken: "asdasdsa",
					},
					{
						Type:          constants.TokenTypeRefresh,
						HashAuthToken: "gfgfgfgf",
					},
				},
			},
			expected: SignInUserRequestResponse{
				User: SignInUserUserRequestResponse{
					UUID: "123",
					Name: "Name",
				},
				AccessToken:  "asdasdsa",
				RefreshToken: "gfgfgfgf",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, ParseSignInUserRequestResponse(tc.user))
		})
	}
}

func TestValidateRefreshUserTokenRequest(t *testing.T) {
	for _, tc := range []struct {
		name                    string
		refreshUserTokenRequest RefreshUserTokenRequest
		expectedError           error
	}{
		{
			name: "ShouldSucessfullyValidateAll",
			refreshUserTokenRequest: RefreshUserTokenRequest{
				RefreshToken: "123",
			},
		},
		{
			name: "ShouldReturnErrorWhenFieldEmpty",
			refreshUserTokenRequest: RefreshUserTokenRequest{
				RefreshToken: "",
			},
			expectedError: errors.New("empty-parameter"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedError, tc.refreshUserTokenRequest.Validate())
		})
	}
}

func TestParseRefreshUserTokenResponse(t *testing.T) {
	for _, tc := range []struct {
		name     string
		user     entity.User
		expected RefreshUserTokenResponse
	}{
		{
			name: "ShouldSucessfullyValidateAll",
			user: entity.User{
				UUID: "123",
				Name: "Name",
				Tokens: []entity.TokenWrapper{
					{
						HashAuthToken: "123",
						Type:          constants.TokenTypeAccess,
					},
					{
						HashAuthToken: "456",
						Type:          constants.TokenTypeRefresh,
					},
				},
			},
			expected: RefreshUserTokenResponse{
				AccessToken:  "123",
				RefreshToken: "456",
				User: RefreshUserTokenUserResponse{
					UUID: "123",
					Name: "Name",
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, ParseRefreshUserTokenResponse(tc.user))
		})
	}
}
