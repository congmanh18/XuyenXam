package handle_product_category

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (h *productCategoryHandlerImpl) HandleGetByProductID(c echo.Context) error {
	productID := c.Param("id")

	categories, err := h.getByProductIDUseCase.Execute(c.Request().Context(), &productID)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Get product categories failed: "+err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", categories)
}
