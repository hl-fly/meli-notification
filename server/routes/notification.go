package routes

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server/handler"
	"github.com/labstack/echo/v4"
)

func registerNotification(app application.App, admin *echo.Group) *echo.Group {
	notification := admin.Group("/notification")
	notification.POST("", handler.HandleInsertNotification(app.Services().Notification()))

	return notification
}
