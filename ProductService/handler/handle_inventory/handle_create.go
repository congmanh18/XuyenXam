package handle_inventory

import (
	"net/http"

	mapper "github.com/congmanh18/XuyenXam/ProductService/mapper/transformentity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"

	"github.com/labstack/echo/v4"
)

func (h *inventoryHandleImpl) HandleCreate(c echo.Context) error {
	var req req.InventoryReq

	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validate error: "+err.Error())
	}

	inventory := mapper.TransformInventoryReqToEntity(req)
	return h.createUseCase.Execute(c.Request().Context(), inventory.ProductID, inventory.Quantity)
}
