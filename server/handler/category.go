package handler

import (
	"net/http"

	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/server/viewmodel"
	"github.com/labstack/echo/v4"
)

func HandleInsertCategory(categoryService contract.CategoryService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var vm viewmodel.CategoryInsertRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		errors := vm.Validate()
		if errors != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors(errors))
		}

		category, err := categoryService.Insert(vm.Parse())
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondJSON(ctx, http.StatusCreated, viewmodel.ParseCategoryInsertResponse(category))
	}
}
