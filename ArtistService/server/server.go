package server

import (
	"github.com/congmanh18/XuyenXam/ArtistService/handler"
	"github.com/congmanh18/XuyenXam/ArtistService/migration"
	"github.com/congmanh18/XuyenXam/ArtistService/model"
	"github.com/congmanh18/XuyenXam/ArtistService/provider"
	"github.com/congmanh18/XuyenXam/ArtistService/repository"
	"github.com/congmanh18/XuyenXam/ArtistService/usecase"
	"github.com/congmanh18/lucas-core/config"
	. "github.com/congmanh18/lucas-core/transport/http"
	"github.com/congmanh18/lucas-core/transport/http/engine"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

const allowMigration = false

func NewServer(serviceConf model.ServiceConfig, routes []route.GroupRoute) *Server {
	var e = engine.NewEcho()
	return NewHttpServer(
		AddName(serviceConf.ServiceName),
		AddPort(serviceConf.ServicePort),
		AddEngine(e),
		AddGroupRoutes(routes),
	)
}

func Run(confPath string) {
	var serviceConf model.ServiceConfig
	config.MustLoadConfig(confPath, &serviceConf)

	var (
		appProvider   = provider.NewAppProvider(serviceConf)
		artistRepo    = repository.NewRepo(appProvider.Postgres)
		artistHandler = handler.NewArtistHandler(handler.ArtistHandlerInject{
			CreateUseCase:   usecase.NewCreateUseCase(artistRepo),
			FindByIDUseCase: usecase.NewFindByIDUseCase(artistRepo),
			FindAllUseCase:  usecase.NewFindAllUseCase(artistRepo),
			UpdateUseCase:   usecase.NewCreateUseCase(artistRepo),
			DeleteUseCase:   usecase.NewDeleteUseCase(artistRepo),
		})
		routes = Routes(artistHandler)
		server = NewServer(serviceConf, routes)
	)

	if allowMigration {
		migration.Must(appProvider.Postgres.Executor)
	}

	server.Run()
}
