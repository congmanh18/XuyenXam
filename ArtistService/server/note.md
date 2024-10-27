package server

import (
	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/XuyenXam/ArtistService/handler"
	"github.com/congmanh18/lucas-core/middleware"
	"github.com/congmanh18/lucas-core/transport/http/method"
	"github.com/congmanh18/lucas-core/transport/http/route"
	"github.com/labstack/echo/v4"
)

// e.PUT("/artist/:id", handler.HandleUpdate, middleware.FindEntityByID(useCase, "artist"),)

func Routes(handler handler.ArtistHandler, useCase middleware.FindByIDUseCase[entity.Artist]) []route.GroupRoute {
	findArtistMiddleware := middleware.FindEntityByID(useCase, "id")
	return []route.GroupRoute{
		{
			Prefix: "/artists",
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: handler.HandleCreate,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: handler.HandleFindByID,
				},
				{
					Path:    "/",
					Method:  method.GET,
					Handler: handler.HandleFindAll,
				},
				{
					Path:    "/delete/:id",
					Method:  method.DELETE,
					Handler: handler.HandleFindByID,
					Middlewares: []echo.MiddlewareFunc{
						findArtistMiddleware,
					},
				},
				{
					Path:    "/update/:id",
					Method:  method.PUT,
					Handler: handler.HandleFindByID,
					Middlewares: []echo.MiddlewareFunc{
						findArtistMiddleware,
					},
				},
			},
		},
	}
}
