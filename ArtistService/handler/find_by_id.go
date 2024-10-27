package handler

import (
	"net/http"

	"github.com/congmanh18/XuyenXam/ArtistService/model/req"
	"github.com/congmanh18/lucas-core/pointer"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

func (a *artistHandlerImpl) HandleFindByID(c echo.Context) error {
	var artist req.BaseReq

	// Bind and validate the request
	if err := c.Bind(&artist); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	if err := validate.Struct(&artist); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validation Error: "+err.Error())
	}

	// Execute use case
	artistEntity, err := a.findByIDUseCase.Execute(c.Request().Context(), pointer.String(artist.ID))
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to find artist: "+err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", artistEntity)
}
