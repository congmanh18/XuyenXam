package handle_category

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleGetByID implements CategoryHandler.
func (h *categoryHandleImpl) HandleGetByID(ctx echo.Context) error {
	id := ctx.Param("id")

	category, err := h.getByIDUseCase.Execute(ctx.Request().Context(), &id)
	if err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "Get category by id failed: "+err.Error())
	}

	return response.OK(ctx, http.StatusOK, "OK", category)
}
