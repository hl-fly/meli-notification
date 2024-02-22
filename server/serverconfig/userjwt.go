package serverconfig

import (
	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/infra/auth"
	"github.com/hector-leite/meli-notification/server/serverutils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserJWTMiddleware(publicKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
			TokenLookup:   "header:" + auth.HTTPHeader,
			ContextKey:    constants.ContextKeyUser,
			SigningKey:    []byte(publicKey),
			SigningMethod: auth.TokenSigningMethod.Name,
			Claims:        new(auth.UserClaims),
		})

		return jwtMiddleware(func(ctx echo.Context) error {
			token := serverutils.GetToken(ctx, constants.ContextKeyUser)
			claims := serverutils.GetUserClaims(token)
			serverutils.GetAccessToken(ctx)

			context := serverutils.SetJWTContext(ctx)
			ctx.Set(constants.ContextKeyDomainUserClaims, claims)
			ctx.Set(constants.ContextKeyDomain, context)

			return next(ctx)
		})
	}
}
