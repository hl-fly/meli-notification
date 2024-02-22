package application

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/service/authtokenservice"
	"github.com/hector-leite/meli-notification/infra/config"
)

func ProvideAuthTokenService(
	data contract.DataManager,
	serverConfig config.ServerAuthConfig,
) contract.AuthTokenService {
	return authtokenservice.NewAuthTokenService(
		data,
		serverConfig.JWTSignatureKey,
	)
}
