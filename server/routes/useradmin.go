package routes

import (
	"github.com/hector-leite/meli-notification/application"
	"github.com/hector-leite/meli-notification/server/serverconfig"
	"github.com/labstack/echo/v4"
)

func registerUserAdmin(app application.App, user *echo.Group) *echo.Group {
	userAdmin := user.Group("/admin", serverconfig.UserAdminJWTMiddleware())

	return userAdmin
}
