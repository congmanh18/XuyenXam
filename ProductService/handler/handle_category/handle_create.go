package handle_category

import (
	"net/http"

	mapper "github.com/congmanh18/XuyenXam/ProductService/mapper/transformentity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleCreate implements CategoryHandler.
func (c *categoryHandleImpl) HandleCreate(ctx echo.Context) error {
	var category req.CategoryReq

	if err := ctx.Bind(&category); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := validate.Struct(&category); err != nil {
		return response.Error(ctx, http.StatusBadRequest, "Validate error: "+err.Error())
	}

	if err := c.createUseCase.Execute(ctx.Request().Context(), mapper.TransformCategoryReqToEntity(category)); err != nil {
		return response.Error(ctx, http.StatusInternalServerError, "Create category failed: "+err.Error())
	}

	return response.OK(ctx, http.StatusOK, "OK", nil)
}
