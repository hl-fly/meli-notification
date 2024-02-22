package handler

import (
	"net/http"

	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/server/viewmodel"
	"github.com/labstack/echo/v4"
)

func HandleSignUp(userService contract.UserService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var vm viewmodel.SignUpUserRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		errors := vm.Validate()
		if errors != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors(errors))
		}

		err = userService.SignUp(vm.Password, vm.Parse())
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondNoContent(ctx)
	}
}

func HandleSignIn(userService contract.UserService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var vm viewmodel.SignInUserRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		errors := vm.Validate()
		if errors != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors(errors))
		}

		user, err := userService.SignIn(vm.Password, vm.Parse())
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondJSON(ctx, http.StatusOK, viewmodel.ParseSignInUserRequestResponse(user))
	}
}

func HandleRefreshToken(userService contract.UserService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var vm viewmodel.RefreshUserTokenRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		err = vm.Validate()
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		user, err := userService.RefreshToken(vm.RefreshToken)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondJSON(ctx, http.StatusOK, viewmodel.ParseRefreshUserTokenResponse(user))
	}
}

func HandleGetUserNotifications(userService contract.UserService, userNotificationService contract.UserNotificationService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		context := getContext(ctx)
		user, err := userService.GetFromContext(context)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		userNotifications, err := userNotificationService.FindByUser(user)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondJSON(ctx, http.StatusOK, viewmodel.ParseGetUserNotificationsResponse(userNotifications))
	}
}
