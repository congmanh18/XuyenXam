package handle_inventory

import (
	"net/http"

	mapper "github.com/congmanh18/XuyenXam/ProductService/mapper/transformentity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (h *inventoryHandleImpl) HandleUpdate(c echo.Context) error {
	id := c.Param("id")

	var req req.InventoryReq

	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validate error: "+err.Error())
	}

	inventory := mapper.TransformInventoryReqToEntity(req)
	inventory.ProductID = &id
	if err := h.updateUseCase.Execute(c.Request().Context(), inventory); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Update inventory failed: "+err.Error())
	}
	return response.OK(c, http.StatusOK, "OK", nil)
}
