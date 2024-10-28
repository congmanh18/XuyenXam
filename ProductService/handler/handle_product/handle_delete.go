package handle_product

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (p *productHandleImpl) HandleDelete(c echo.Context) error {
	id := c.Param("id")

	if err := p.deleteUseCase.Execute(c.Request().Context(), &id); err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", nil)
}
