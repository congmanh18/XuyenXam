package handle_category

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleGetProductByCategory implements CategoryHandler.
func (h *categoryHandleImpl) HandleGetProductByCategory(c echo.Context) error {
	categoryID := c.Param("id")

	products, err := h.getProductByCategoryUseCase.Execute(c.Request().Context(), &categoryID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Get products by category failed: "+err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", products)
}
