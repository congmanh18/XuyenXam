package server

import (
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_category"
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_image"
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_inventory"
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_product"
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_product_category"
	"github.com/congmanh18/lucas-core/transport/http/method"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

func Routes(
	product handle_product.ProductHandler,
	category handle_category.CategoryHandler,
	product_category handle_product_category.ProductCategoryHandler,
	image handle_image.ImageHandler,
	inventory handle_inventory.InventoryHandler,
) []route.GroupRoute {
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
				{
					Path:    "/paginate",
					Method:  method.GET,
					Handler: product.HandleGetPaginate,
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
					Path:    "/product-list",
					Method:  method.GET,
					Handler: category.HandleGetProductByCategory,
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
		{
			Prefix: "/product-category",
			Routes: []route.Route{
				{
					Path:    "/",
					Method:  method.POST,
					Handler: product_category.HandleAssign,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: product_category.HandleGetByProductID,
				},
				{
					Path:    "/remove",
					Method:  method.DELETE,
					Handler: product_category.HandleRemove,
				},
			},
		},
		{
			Prefix: "/product-image",
			Routes: []route.Route{
				{
					Path:    "/",
					Method:  method.POST,
					Handler: image.HandleAddImage,
				},
				{
					Path:    "/:id",
					Method:  method.DELETE,
					Handler: image.HandleDeleteImage,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: image.HandleGetImageByProductID,
				},
			},
		},
		{
			Prefix: "/product-inventory",
			Routes: []route.Route{
				{
					Path:    "/",
					Method:  method.POST,
					Handler: inventory.HandleCreate,
				},
				{
					Path:    "/:id",
					Method:  method.GET,
					Handler: inventory.HandleGetByProductID,
				},
				{
					Path:    "/update/:id",
					Method:  method.PUT,
					Handler: inventory.HandleUpdate,
				},
				{
					Path:    "/delete/:id",
					Method:  method.DELETE,
					Handler: inventory.HandleDelete,
				},
				{
					Path:    "/bulk-update",
					Method:  method.PUT,
					Handler: inventory.HandleBulkUpdate,
				},
			},
		},
	}
}
