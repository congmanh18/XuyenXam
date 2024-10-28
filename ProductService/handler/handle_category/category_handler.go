package handle_category

import (
	"github.com/congmanh18/XuyenXam/ProductService/usecase/category_usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type CategoryHandler interface {
	HandleCreate(c echo.Context) error
	HandleGetByID(c echo.Context) error
	HandleGetAll(c echo.Context) error
	HandleUpdate(c echo.Context) error
	HandleDelete(c echo.Context) error
}

type categoryHandleImpl struct {
	createUseCase  category_usecase.CreateUseCase
	getByIDUseCase category_usecase.GetByIDUseCase
	getAllUseCase  category_usecase.GetAllUseCase
	updateUseCase  category_usecase.UpdateUseCase
	deleteUseCase  category_usecase.DeleteUseCase
}

type CategoryHandleInject struct {
	CreateUseCase  category_usecase.CreateUseCase
	GetByIDUseCase category_usecase.GetByIDUseCase
	GetAllUseCase  category_usecase.GetAllUseCase
	UpdateUseCase  category_usecase.UpdateUseCase
	DeleteUseCase  category_usecase.DeleteUseCase
}

func NewCategoryHandler(inject CategoryHandleInject) CategoryHandler {
	return &categoryHandleImpl{
		createUseCase:  inject.CreateUseCase,
		getByIDUseCase: inject.GetByIDUseCase,
		getAllUseCase:  inject.GetAllUseCase,
		updateUseCase:  inject.UpdateUseCase,
		deleteUseCase:  inject.DeleteUseCase,
	}
}
