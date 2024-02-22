package serverconfig

import (
	"errors"

	"github.com/hector-leite/meli-notification/server/serverutils"
	"github.com/labstack/echo/v4"
)

func UserAdminJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			claims := serverutils.GetUserClaimsFromContextDomain(ctx)
			if claims.Role == "admin" {
				return next(ctx)
			}
			return errors.New("Unauthorized")
		}
	}
}
