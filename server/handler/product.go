package handler

import (
	"net/http"

	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/server/viewmodel"
	"github.com/labstack/echo/v4"
)

func HandleInsertProduct(productService contract.ProductService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var vm viewmodel.ProductInsertRequest
		err := ctx.Bind(&vm)
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		errors := vm.Validate()
		if errors != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors(errors))
		}

		product, err := productService.Insert(vm.Parse())
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondJSON(ctx, http.StatusCreated, viewmodel.ParseProductInsertResponse(product))
	}
}

func HandleGetProducts(productService contract.ProductService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		products, err := productService.Find()
		if err != nil {
			return respondJSON(ctx, http.StatusUnprocessableEntity, viewmodel.FormatErrors([]error{err}))
		}

		return respondJSON(ctx, http.StatusOK, viewmodel.ParseGetProductsResponse(products))
	}
}
