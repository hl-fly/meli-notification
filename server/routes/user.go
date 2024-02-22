package routes

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server/handler"
	"github.com/hector-leite/meli-notification/server/serverconfig"
	"github.com/labstack/echo/v4"
)

func registerUser(app application.App, root *echo.Group) *echo.Group {
	user := root.Group("/user", serverconfig.UserJWTMiddleware(app.Config.Server.Auth.JWTSignatureKey))

	user.GET("/notifications", handler.HandleGetUserNotifications(app.Services().User(), app.Services().UserNotification()))

	return user
}
