package application

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/infra/config"
)

type App struct {
	Config       *config.Config
	DataManager  contract.DataManager
	CacheManager contract.CacheManager
	HTTPClient   contract.HTTPClient
	services     *AppService
}

func BuildApp() App {
	return injectApp()
}

func (app App) Services() *AppService {
	return app.services
}
