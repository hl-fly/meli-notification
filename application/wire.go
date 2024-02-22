//go:build wireinject
// +build wireinject

package application

import (
	"github.com/google/wire"
	"github.com/hector-leite/meli-notification/domain/service/categoryservice"
	"github.com/hector-leite/meli-notification/domain/service/notificationservice"
	"github.com/hector-leite/meli-notification/domain/service/productservice"
	"github.com/hector-leite/meli-notification/domain/service/usernotificationservice"
	"github.com/hector-leite/meli-notification/domain/service/userproducthistoryservice"
	"github.com/hector-leite/meli-notification/domain/service/userservice"
	"github.com/hector-leite/meli-notification/infra/config"
)

// injectApp builds App
func injectApp() App {
	panic(wire.Build(
		AppSet,

		// If you add any field added to App struct, you MUST add it here so that wire can correctly
		// inject it.
		wire.Struct(new(App),
			"Config",
			"DataManager",
			"CacheManager",
			"HTTPClient",
			"services",
		),
		wire.FieldsOf(new(*config.ServerConfig),
			"Auth",
		),
	))
}

var AppSet = wire.NewSet(
	bindingSet,

	ProvideDB,
	ProvideCacheManager,
	ProvideConfig,
	ProvideHTTPClient,

	NewAppService,
)

var bindingSet = wire.NewSet(
	configSet,
	serviceSet,
)

var configSet = wire.NewSet(
	wire.FieldsOf(new(*config.Config),
		"App",
		"Cache",
		"DB",
		"Server",
	),
)
var serviceSet = wire.NewSet(
	ProvideAuthTokenService,
	userservice.NewUserService,
	categoryservice.NewCategoryService,
	productservice.NewProductService,
	userproducthistoryservice.NewUserProductHistoryService,
	notificationservice.NewNotificationService,
	usernotificationservice.NewUserNotificationService,
)
