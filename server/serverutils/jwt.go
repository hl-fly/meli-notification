package serverutils

import (
	"github.com/golang-jwt/jwt"
	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/infra/auth"
	"github.com/hector-leite/meli-notification/infra/context"
	"github.com/labstack/echo/v4"
)

func GetToken(ctx echo.Context, contextKey string) *jwt.Token {
	return ctx.Get(contextKey).(*jwt.Token)
}

func GetUserClaims(token *jwt.Token) *auth.UserClaims {
	return token.Claims.(*auth.UserClaims)
}

func SetJWTContext(ctx echo.Context) context.Context {
	context := GetContextDomain(ctx)

	rawToken := ctx.Get(constants.ContextKeyUser)
	if rawToken != nil {
		token := rawToken.(*jwt.Token)
		claims := token.Claims.(*auth.UserClaims)
		context.Set(constants.ContextKeyUserUUID, claims.UUID)
	}

	rawUserUUID := ctx.Get(constants.ContextKeyUserUUID)
	if rawUserUUID != nil {
		context.Set(constants.ContextKeyUserUUID, rawUserUUID.(string))
	}

	return context
}
