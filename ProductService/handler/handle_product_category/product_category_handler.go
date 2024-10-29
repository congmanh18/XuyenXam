package handle_product_category

import (
	usecase "github.com/congmanh18/XuyenXam/ProductService/usecase/product_category_usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type ProductCategoryHandler interface {
	HandleAssign(c echo.Context) error
	HandleGetByProductID(c echo.Context) error
	HandleRemove(c echo.Context) error
}

type productCategoryHandlerImpl struct {
	assignUseCase         usecase.AssignUseCase
	getByProductIDUseCase usecase.GetByProductIDUseCase
	removeUseCase         usecase.RemoveUseCase
}

type ProductCategoryHandlerInject struct {
	AssignUseCase         usecase.AssignUseCase
	GetByProductIDUseCase usecase.GetByProductIDUseCase
	RemoveUseCase         usecase.RemoveUseCase
}

func NewProductCategoryHandler(inject ProductCategoryHandlerInject) ProductCategoryHandler {
	return &productCategoryHandlerImpl{assignUseCase: inject.AssignUseCase, getByProductIDUseCase: inject.GetByProductIDUseCase, removeUseCase: inject.RemoveUseCase}
}
