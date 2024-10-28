package server

import (
	"context"
	"log"
	"time"

	"github.com/congmanh18/XuyenXam/ArtistService/handler"
	"github.com/congmanh18/XuyenXam/ArtistService/migration"
	"github.com/congmanh18/XuyenXam/ArtistService/model"
	"github.com/congmanh18/XuyenXam/ArtistService/provider"
	"github.com/congmanh18/XuyenXam/ArtistService/repository"
	"github.com/congmanh18/XuyenXam/ArtistService/usecase"
	"github.com/congmanh18/lucas-core/config"
	. "github.com/congmanh18/lucas-core/transport/http"
)

const allowMigration = false

func Run(confPath string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var serviceConf model.ServiceConfig
	config.MustLoadConfig(confPath, &serviceConf)

	appProvider := provider.NewAppProvider(serviceConf)
	artistRepo := repository.NewRepo(appProvider.MongoDB)
	artistHandler := buildArtistHandler(artistRepo)
	routes := Routes(artistHandler)

	server := NewServer
	// NewServer(serviceConf, routes)

	if allowMigration {
		if err := migration.Must(ctx, appProvider.MongoDB.Client); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	server.Run()
}

func buildArtistHandler(repo repository.Repo) *handler.ArtistHandler {
	return handler.NewArtistHandler(handler.ArtistHandlerInject{
		CreateUseCase:   usecase.NewCreateUseCase(repo),
		FindByIDUseCase: usecase.NewFindByIDUseCase(repo),
		FindAllUseCase:  usecase.NewFindAllUseCase(repo),
		UpdateUseCase:   usecase.NewUpdateUseCase(repo),
		DeleteUseCase:   usecase.NewDeleteUseCase(repo),
	})
}
