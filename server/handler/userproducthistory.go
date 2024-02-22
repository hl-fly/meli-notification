package handler

import (
	"net/http"

	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/server/viewmodel"
	"github.com/labstack/echo/v4"
)

func HandleInsertUserProductHistory(userService contract.UserService,
	userProductHistoryService contract.UserProductHistoryService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var vm viewmodel.UserProductHistoryRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		errors := vm.Validate()
		if errors != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors(errors))
		}

		context := getContext(ctx)
		user, err := userService.GetFromContext(context)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors(errors))
		}

		err = userProductHistoryService.Insert(vm.Parse(user))
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondNoContent(ctx)
	}
}
