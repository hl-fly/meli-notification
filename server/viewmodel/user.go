package viewmodel

import (
	"errors"
	"fmt"

	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/domain/entity"
)

type SignUpUserRequest struct {
	Name            string `json:"name"`
	Password        string `json:"password"`
	CPF             string `json:"cpf"`
	Email           string `json:"email"`
	ConfirmPassword string `json:"confirm_password"`
}

func (vm SignUpUserRequest) Validate() []error {
	genericErrorFormat := "invalid-parameter-%s"
	var aggErrors []error
	if vm.Name == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "name"))
	}

	if vm.Email == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "email"))
	}

	if vm.CPF == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "cpf"))
	}

	if vm.Password == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "password"))
	}

	if vm.ConfirmPassword == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "confirm-password"))
	}

	if (vm.Password != "" && vm.ConfirmPassword != "") && (vm.Password != vm.ConfirmPassword) {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "password-confirm-password"))
	}

	if len(aggErrors) > 0 {
		return aggErrors
	}

	return nil
}

func (vm SignUpUserRequest) Parse() entity.User {
	return entity.User{
		Name:  vm.Name,
		Email: vm.Email,
		CPF:   vm.CPF,
	}
}

type SignInUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (vm SignInUserRequest) Validate() []error {
	genericErrorFormat := "invalid-parameter-%s"
	var aggErrors []error

	if vm.Email == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "email"))
	}

	if vm.Password == "" {
		aggErrors = append(aggErrors, fmt.Errorf(genericErrorFormat, "password"))
	}

	if len(aggErrors) > 0 {
		return aggErrors
	}

	return nil
}

func (vm SignInUserRequest) Parse() entity.User {
	return entity.User{
		Email: vm.Email,
	}
}

type SignInUserRequestResponse struct {
	RefreshToken string                        `json:"refresh_token,omitempty"`
	AccessToken  string                        `json:"access_token,omitempty"`
	User         SignInUserUserRequestResponse `json:"user"`
}

type SignInUserUserRequestResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func ParseSignInUserRequestResponse(user entity.User) SignInUserRequestResponse {
	var vm SignInUserRequestResponse
	vm.User.UUID = user.UUID
	vm.User.Name = user.Name

	for _, token := range user.Tokens {
		switch token.Type {
		case constants.TokenTypeAccess:
			vm.AccessToken = token.HashAuthToken
		case constants.TokenTypeRefresh:
			vm.RefreshToken = token.HashAuthToken
		}
	}

	return vm
}

type RefreshUserTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (p RefreshUserTokenRequest) Validate() error {
	if p.RefreshToken == "" {
		return errors.New("empty-parameter")
	}

	return nil
}

type RefreshUserTokenResponse struct {
	RefreshToken string                       `json:"refresh_token,omitempty"`
	AccessToken  string                       `json:"access_token,omitempty"`
	User         RefreshUserTokenUserResponse `json:"user"`
}

type RefreshUserTokenUserResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func ParseRefreshUserTokenResponse(user entity.User) RefreshUserTokenResponse {
	var vm RefreshUserTokenResponse
	vm.User.UUID = user.UUID
	vm.User.Name = user.Name

	for _, token := range user.Tokens {
		switch token.Type {
		case constants.TokenTypeAccess:
			vm.AccessToken = token.HashAuthToken
		case constants.TokenTypeRefresh:
			vm.RefreshToken = token.HashAuthToken
		}
	}

	return vm
}

type GetUserNotificationsResponse []GetUserNotificationResponse

type GetUserNotificationResponse struct {
	Message string
	Link    string
}

func ParseGetUserNotificationsResponse(userNotifications []entity.UserNotification) GetUserNotificationsResponse {
	userNotificationsResponse := make([]GetUserNotificationResponse, len(userNotifications))
	for i := 0; i < len(userNotifications); i++ {
		userNotificationsResponse[i].Message = userNotifications[i].Notification.Message
		userNotificationsResponse[i].Link = userNotifications[i].Notification.Link
	}

	return userNotificationsResponse
}
