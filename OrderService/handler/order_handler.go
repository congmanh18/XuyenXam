package handler

import (
	"github.com/congmanh18/XuyenXam/OrderService/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type OrderHandler interface {
	HandleCreate(c echo.Context) error
	HandleGetByID(c echo.Context) error
	HanldeGetAll(c echo.Context) error
	HandleUpdate(c echo.Context) error
	HandleDelete(c echo.Context) error
}

type orderHandleImpl struct {
	createUseCase  usecase.CreateUseCase
	getByIDUseCase usecase.GetByIDUseCase
	getAllUseCase  usecase.GetAllUseCase
	updateUseCase  usecase.UpdateUseCase
	deleteUseCase  usecase.DeleteUseCase
}

type OrderHandleInject struct {
	CreateUseCase  usecase.CreateUseCase
	GetByIDUseCase usecase.GetByIDUseCase
	GetAllUseCase  usecase.GetAllUseCase
	UpdateUseCase  usecase.UpdateUseCase
	DeleteUseCase  usecase.DeleteUseCase
}

func NewOrderHandler(inj OrderHandleInject) OrderHandler {
	return &orderHandleImpl{
		createUseCase:  inj.CreateUseCase,
		getByIDUseCase: inj.GetByIDUseCase,
		getAllUseCase:  inj.GetAllUseCase,
		updateUseCase:  inj.UpdateUseCase,
		deleteUseCase:  inj.DeleteUseCase,
	}
}
