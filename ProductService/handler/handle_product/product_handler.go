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
	HandleUpdate(c echo.Context) error
	HandleDelete(c echo.Context) error
}

type productHandleImpl struct {
	createUseCase  product_usecase.CreateUseCase
	getByIDUseCase product_usecase.GetByIDUseCase
	getAllUseCase  product_usecase.GetAllUseCase
	updateUseCase  product_usecase.UpdateUseCase
	deleteUseCase  product_usecase.DeleteUseCase
}

type ProducHandleInject struct {
	CreateUseCase  product_usecase.CreateUseCase
	GetByIDUseCase product_usecase.GetByIDUseCase
	GetAllUseCase  product_usecase.GetAllUseCase
	UpdateUseCase  product_usecase.UpdateUseCase
	DeleteUseCase  product_usecase.DeleteUseCase
}

func NewProductHandler(inj ProducHandleInject) ProductHandler {
	return &productHandleImpl{
		createUseCase:  inj.CreateUseCase,
		getByIDUseCase: inj.GetByIDUseCase,
		getAllUseCase:  inj.GetAllUseCase,
		updateUseCase:  inj.UpdateUseCase,
		deleteUseCase:  inj.DeleteUseCase,
	}
}
