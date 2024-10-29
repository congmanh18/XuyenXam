package handle_product

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/congmanh18/XuyenXam/ProductService/usecase/product_usecase"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type ProductHandler interface {
	HandleCreate(c echo.Context) error
	HandleGetByID(c echo.Context) error
	HandleGetAll(c echo.Context) error
	HandleGetPaginate(c echo.Context) error
	HandleUpdate(c echo.Context) error
	HandleDelete(c echo.Context) error
	HandleSearch(c echo.Context) error
}

type productHandleImpl struct {
	createUseCase      product_usecase.CreateUseCase
	getByIDUseCase     product_usecase.GetByIDUseCase
	getAllUseCase      product_usecase.GetAllUseCase
	getPaginateUseCase product_usecase.GetPaginateUseCase
	updateUseCase      product_usecase.UpdateUseCase
	deleteUseCase      product_usecase.DeleteUseCase
	searchUseCase      product_usecase.SearchUseCase
}

type ProducHandleInject struct {
	CreateUseCase      product_usecase.CreateUseCase
	GetByIDUseCase     product_usecase.GetByIDUseCase
	GetAllUseCase      product_usecase.GetAllUseCase
	GetPaginateUseCase product_usecase.GetPaginateUseCase
	UpdateUseCase      product_usecase.UpdateUseCase
	DeleteUseCase      product_usecase.DeleteUseCase
	SearchUseCase      product_usecase.SearchUseCase
}

func NewProductHandler(inj ProducHandleInject) ProductHandler {
	return &productHandleImpl{
		createUseCase:      inj.CreateUseCase,
		getByIDUseCase:     inj.GetByIDUseCase,
		getAllUseCase:      inj.GetAllUseCase,
		getPaginateUseCase: inj.GetPaginateUseCase,
		updateUseCase:      inj.UpdateUseCase,
		deleteUseCase:      inj.DeleteUseCase,
		searchUseCase:      inj.SearchUseCase,
	}
}
