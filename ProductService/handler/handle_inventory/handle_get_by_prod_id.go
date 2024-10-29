package handle_inventory

import (
	"net/http"

	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (h *inventoryHandleImpl) HandleGetByProductID(c echo.Context) error {
	id := c.Param("id")

	var req req.InventoryReq

	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validate error: "+err.Error())
	}

	inventory, err := h.getByProductIDUseCase.Execute(c.Request().Context(), &id)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Get inventory failed: "+err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", inventory)
}
