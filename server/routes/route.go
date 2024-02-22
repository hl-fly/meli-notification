package routes

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server/handler"
	"github.com/labstack/echo/v4"
)

func Register(app application.App, server *echo.Echo) {
	root := server.Group("/api")

	root.POST("/signup", handler.HandleSignUp(app.Services().User()))
	root.POST("/signin", handler.HandleSignIn(app.Services().User()))
	root.POST("/refresh-token", handler.HandleRefreshToken(app.Services().User()))

	user := registerUser(app, root)
	admin := registerUserAdmin(app, user)
	registerCategory(app, admin)
	registerProductAdm(app, admin)
	registerProductUser(app, user)
	registerUserProductHistory(app, user)
	registerNotification(app, admin)
}
