package handle_category

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleGetAll implements CategoryHandler.
func (h *categoryHandleImpl) HandleGetAll(ctx echo.Context) error {
	category, err := h.getAllUseCase.Execute(ctx.Request().Context())
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "Get all category failed: "+err.Error())
	}

	return response.OK(ctx, http.StatusOK, "OK", category)
}
