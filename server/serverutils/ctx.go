package serverutils

import (
	"strings"

	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/infra/auth"
	"github.com/hector-leite/meli-notification/infra/context"
	"github.com/labstack/echo/v4"
)

func GetAccessToken(ctx echo.Context) string {
	header := ctx.Request().Header.Get(auth.HTTPHeader)
	accessToken := strings.Replace(header, "Bearer ", "", 1)
	return accessToken
}

func GetContextDomain(ctx echo.Context) context.Context {
	ct := ctx.Get(constants.ContextKeyDomain)
	return ct.(context.Context)
}

func GetUserClaimsFromContextDomain(ctx echo.Context) *auth.UserClaims {
	ct := ctx.Get(constants.ContextKeyDomainUserClaims)
	return ct.(*auth.UserClaims)
}
