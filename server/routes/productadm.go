package routes

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server/handler"
	"github.com/labstack/echo/v4"
)

func registerProductAdm(app application.App, admin *echo.Group) *echo.Group {
	product := admin.Group("/product")
	product.POST("", handler.HandleInsertProduct(app.Services().Product()))

	return product
}
