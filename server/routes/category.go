package routes

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server/handler"
	"github.com/labstack/echo/v4"
)

func registerCategory(app application.App, admin *echo.Group) *echo.Group {
	category := admin.Group("/category")
	category.POST("", handler.HandleInsertCategory(app.Services().Category()))

	return category
}
