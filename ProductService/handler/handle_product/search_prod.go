package handle_product

import (
	"net/http"

	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (p *productHandleImpl) HandleSearch(c echo.Context) error {
	query := c.QueryParam("query")
	products, err := p.searchUseCase.Execute(c.Request().Context(), &query)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}
	return response.OK(c, http.StatusOK, "OK", products)
}
