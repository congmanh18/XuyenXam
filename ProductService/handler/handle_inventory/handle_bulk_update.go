package handle_inventory

import (
	"net/http"

	"github.com/congmanh18/XuyenXam/ProductService/entity"
	mapper "github.com/congmanh18/XuyenXam/ProductService/mapper/transformentity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (h *inventoryHandleImpl) HandleBulkUpdate(c echo.Context) error {
	var req []req.InventoryReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	// check sau, chưa biết code đúng không
	err := validate.Struct(&req)
	if len(req) == 0 || err != nil {
		return response.Error(c, http.StatusBadRequest, "Request body is empty")
	}

	inventories := make([]entity.Inventory, 0, len(req))
	for _, inventory := range req {
		inventories = append(inventories, *mapper.TransformInventoryReqToEntity(inventory))
	}

	if err := h.bulkUpdateUseCase.Execute(c.Request().Context(), inventories); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Bulk update inventory failed: "+err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", nil)
}
