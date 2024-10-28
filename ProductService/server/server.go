package server

import (
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_category"
	"github.com/congmanh18/XuyenXam/ProductService/handler/handle_product"
	"github.com/congmanh18/XuyenXam/ProductService/migration"
	"github.com/congmanh18/XuyenXam/ProductService/model"
	"github.com/congmanh18/XuyenXam/ProductService/provider"
	"github.com/congmanh18/XuyenXam/ProductService/repository/category"
	"github.com/congmanh18/XuyenXam/ProductService/repository/product"
	"github.com/congmanh18/XuyenXam/ProductService/usecase/category_usecase"
	"github.com/congmanh18/XuyenXam/ProductService/usecase/product_usecase"
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
		appProvider = provider.NewAppProvider(serviceConf)

		productRepo  = product.NewProductRepository(appProvider.Postgres)
		categoryRepo = category.NewCategoryRepository(appProvider.Postgres)

		productHandler = handle_product.NewProductHandler(handle_product.ProducHandleInject{
			CreateUseCase:  product_usecase.NewCreateUseCase(productRepo),
			GetByIDUseCase: product_usecase.NewGetByIDUseCase(productRepo),
			GetAllUseCase:  product_usecase.NewGetAllUseCase(productRepo),
			UpdateUseCase:  product_usecase.NewUpdateUseCase(productRepo),
			DeleteUseCase:  product_usecase.NewDeleteUseCase(productRepo),
		})

		categoryHandler = handle_category.NewCategoryHandler(handle_category.CategoryHandleInject{
			CreateUseCase:  category_usecase.NewCreateUseCase(categoryRepo),
			GetByIDUseCase: category_usecase.NewGetByIDUseCase(categoryRepo),
			GetAllUseCase:  category_usecase.NewGetAllUseCase(categoryRepo),
			UpdateUseCase:  category_usecase.NewUpdateUseCase(categoryRepo),
			DeleteUseCase:  category_usecase.NewDeleteUseCase(categoryRepo),
		})

		routes = Routes(productHandler, categoryHandler)
		server = NewServer(serviceConf, routes)
	)

	if allowMigration {
		migration.Must(appProvider.Postgres.Executor)
	}

	server.Run()
}
