package handler

import (
	"net/http"

	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/server/viewmodel"
	"github.com/labstack/echo/v4"
)

func HandleInsertNotification(notificationService contract.NotificationService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var vm viewmodel.NotificationInsertRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		errors := vm.Validate()
		if errors != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors(errors))
		}

		err = notificationService.Insert(vm.Parse())
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondNoContent(ctx)
	}
}
