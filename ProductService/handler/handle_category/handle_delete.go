package handle_category

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleDelete implements CategoryHandler.
func (h *categoryHandleImpl) HandleDelete(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := h.deleteUseCase.Execute(ctx.Request().Context(), &id); err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "Delete category failed: "+err.Error())
	}

	return response.OK(ctx, http.StatusOK, "OK", nil)
}
