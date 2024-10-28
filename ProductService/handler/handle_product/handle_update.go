package handle_product

import (
	"net/http"

	mapper "github.com/congmanh18/XuyenXam/ProductService/mapper/transformentity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"

	"github.com/labstack/echo/v4"
)

func (p *productHandleImpl) HandleUpdate(c echo.Context) error {
	id := c.Param("id")

	var product req.ProductUpdateReq
	if err := c.Bind(&product); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := p.updateUseCase.Execute(c.Request().Context(), mapper.TransformProductUpdateReqToEntity(product), &id); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", nil)
}
