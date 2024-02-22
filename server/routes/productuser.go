package routes

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server/handler"
	"github.com/labstack/echo/v4"
)

func registerProductUser(app application.App, user *echo.Group) *echo.Group {
	productUser := user.Group("/product")

	productUser.GET("", handler.HandleGetProducts(app.Services().Product()))

	return productUser
}
