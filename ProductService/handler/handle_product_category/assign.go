package handle_product_category

import (
	"net/http"

	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (h *productCategoryHandlerImpl) HandleAssign(c echo.Context) error {
	var req req.ProductCategoryReq
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validate error: "+err.Error())
	}

	if err := h.assignUseCase.Execute(c.Request().Context(), req.ProductID, req.CategoryID); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Assign product category failed: "+err.Error())
	}
	return response.OK(c, http.StatusOK, "OK", nil)
}
