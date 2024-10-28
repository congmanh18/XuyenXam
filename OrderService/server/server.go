package server

import (
	"github.com/congmanh18/XuyenXam/OrderService/handler"
	"github.com/congmanh18/XuyenXam/OrderService/migration"
	"github.com/congmanh18/XuyenXam/OrderService/model"
	"github.com/congmanh18/XuyenXam/OrderService/usecase"

	"github.com/congmanh18/XuyenXam/OrderService/provider"
	"github.com/congmanh18/XuyenXam/OrderService/repository"
	"github.com/congmanh18/lucas-core/config"

	"github.com/congmanh18/lucas-core/transport/http"
	"github.com/congmanh18/lucas-core/transport/http/engine"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

const allowMigration = true

func NewServer(serviceConf model.ServiceConfig, routes []route.GroupRoute) *http.Server {
	var e = engine.NewEcho()
	return http.NewHttpServer(
		http.AddName(serviceConf.ServiceName),
		http.AddPort(serviceConf.ServicePort),
		http.AddEngine(e),
		http.AddGroupRoutes(routes),
	)
}

func Run(confPath string) {
	var serviceConf model.ServiceConfig
	config.MustLoadConfig(confPath, &serviceConf)

	var (
		appProvider  = provider.NewAppProvider(serviceConf)
		orderRepo    = repository.NewOrderRepository(appProvider.Postgres)
		orderHandler = handler.NewOrderHandler(handler.OrderHandleInject{
			CreateUseCase: usecase.NewCreateUseCase(orderRepo),
		})
		routes = Routes(orderHandler)
		server = NewServer(serviceConf, routes)
	)

	if allowMigration {
		migration.Must(appProvider.Postgres.Executor)
	}

	server.Run()
}
