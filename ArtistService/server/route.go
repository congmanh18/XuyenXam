package server

import (
	"github.com/congmanh18/XuyenXam/ArtistService/handler"
	"github.com/congmanh18/lucas-core/transport/http/method"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

func Routes(handler handler.ArtistHandler) []route.GroupRoute {
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
				},
				{
					Path:    "/update/:id",
					Method:  method.PUT,
					Handler: handler.HandleFindByID,
				},
			},
		},
	}
}
