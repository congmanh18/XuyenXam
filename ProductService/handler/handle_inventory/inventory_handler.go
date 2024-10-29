package handle_inventory

import (
	"github.com/congmanh18/XuyenXam/ProductService/usecase/inventory_usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type InventoryHandler interface {
	HandleCreate(c echo.Context) error
	HandleUpdate(c echo.Context) error
	HandleDelete(c echo.Context) error
	HandleGetByProductID(c echo.Context) error
	HandleBulkUpdate(c echo.Context) error
}

type inventoryHandleImpl struct {
	createUseCase         inventory_usecase.CreateInventoryUseCase
	updateUseCase         inventory_usecase.UpdateInventoryUseCase
	deleteUseCase         inventory_usecase.DeleteInventoryUseCase
	getByProductIDUseCase inventory_usecase.GetByProductIDUseCase
	bulkUpdateUseCase     inventory_usecase.BulkUpdateInventoryUseCase
}

type InventoryHandleInject struct {
	CreateUseCase         inventory_usecase.CreateInventoryUseCase
	UpdateUseCase         inventory_usecase.UpdateInventoryUseCase
	DeleteUseCase         inventory_usecase.DeleteInventoryUseCase
	GetByProductIDUseCase inventory_usecase.GetByProductIDUseCase
	BulkUpdateUseCase     inventory_usecase.BulkUpdateInventoryUseCase
}

func NewInventoryHandler(inject InventoryHandleInject) InventoryHandler {
	return &inventoryHandleImpl{
		createUseCase:         inject.CreateUseCase,
		updateUseCase:         inject.UpdateUseCase,
		deleteUseCase:         inject.DeleteUseCase,
		getByProductIDUseCase: inject.GetByProductIDUseCase,
		bulkUpdateUseCase:     inject.BulkUpdateUseCase,
	}
}
