package routes

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server/handler"
	"github.com/labstack/echo/v4"
)

func registerUserProductHistory(app application.App, user *echo.Group) *echo.Group {
	product := user.Group("/user-product-history")
	product.POST("", handler.HandleInsertUserProductHistory(app.Services().User(), app.Services().UserProductHistory()))

	return product
}
