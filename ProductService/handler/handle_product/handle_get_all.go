package handle_product

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (p *productHandleImpl) HandleGetAll(c echo.Context) error {
	products, err := p.getAllUseCase.Execute(c.Request().Context())
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", products)
}
