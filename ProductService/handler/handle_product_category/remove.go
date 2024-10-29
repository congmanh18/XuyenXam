package handle_product_category

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (h *productCategoryHandlerImpl) HandleRemove(c echo.Context) error {
	productID := c.Param("id")
	categoryID := c.QueryParam("category_id")

	if err := h.removeUseCase.Execute(c.Request().Context(), &productID, &categoryID); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Remove product category failed: "+err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", nil)
}
