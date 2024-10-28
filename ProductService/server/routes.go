package server

import (
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_category"
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_product"
	"github.com/congmanh18/lucas-core/transport/http/method"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

func Routes(product handle_product.ProductHandler, category handle_category.CategoryHandler) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/product",
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: product.HandleCreate,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: product.HandleGetByID,
				},
				{
					Path:    "/",
					Method:  method.GET,
					Handler: product.HandleGetAll,
				},
				{
					Path:    "/update/:id",
					Method:  method.PUT,
					Handler: product.HandleUpdate,
				},
				{
					Path:    "/delete/:id",
					Method:  method.DELETE,
					Handler: product.HandleDelete,
				},
			},
		},
		{
			Prefix: "/category",
			Routes: []route.Route{
				{
					Path:    "/create",
					Method:  method.POST,
					Handler: category.HandleCreate,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: category.HandleGetByID,
				},
				{
					Path:    "/",
					Method:  method.GET,
					Handler: category.HandleGetAll,
				},
				{
					Path:    "/update/:id",
					Method:  method.PUT,
					Handler: category.HandleUpdate,
				},
				{
					Path:    "/delete/:id",
					Method:  method.DELETE,
					Handler: category.HandleDelete,
				},
			},
		},
	}
}
