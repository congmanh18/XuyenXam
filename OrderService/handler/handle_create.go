package handler

import (
	"net/http"

	mapper "github.com/congmanh18/XuyenXam/OrderService/mapper/transformentity"
	"github.com/congmanh18/XuyenXam/OrderService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleCreate implements OrderHandler.
func (o *orderHandleImpl) HandleCreate(c echo.Context) error {
	var order req.OrderCreateReq

	if err := c.Bind(&order); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := validate.Struct(&order); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validate error: "+err.Error())
	}

	if err := o.createUseCase.Execute(c.Request().Context(), mapper.TransformOrderCreateReqToEntity(order)); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Create order failed: "+err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", nil)
}
