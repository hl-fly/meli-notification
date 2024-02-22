package serverconfig

import (
	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/infra/context"
	"github.com/labstack/echo/v4"
)

func DomainContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			context := context.New()
			c.Set(constants.ContextKeyDomain, context)

			return next(c)
		}
	}
}
