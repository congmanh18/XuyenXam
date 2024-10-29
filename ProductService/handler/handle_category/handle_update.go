package handle_category

import (
	"net/http"

	mapper "github.com/congmanh18/XuyenXam/ProductService/mapper/transformentity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleUpdate implements CategoryHandler.
func (h *categoryHandleImpl) HandleUpdate(ctx echo.Context) error {
	id := ctx.Param("id")

	var req req.CategoryReq

	if err := ctx.Bind(&req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := validate.Struct(&req); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "Validate error: "+err.Error())
	}

	if err := h.updateUseCase.Execute(ctx.Request().Context(), mapper.TransformCategoryReqToEntity(req), &id); err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "Update category failed: "+err.Error())
	}

	return response.OK(ctx, http.StatusOK, "OK", nil)
}
