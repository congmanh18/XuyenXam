package handle_inventory

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (h *inventoryHandleImpl) HandleDelete(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := h.deleteUseCase.Execute(ctx.Request().Context(), &id); err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "Delete inventory failed: "+err.Error())
	}

	return response.OK(ctx, http.StatusOK, "OK", nil)
}
