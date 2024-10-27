package handler

import (
	"net/http"

	"github.com/congmanh18/XuyenXam/ArtistService/mapper"
	"github.com/congmanh18/XuyenXam/ArtistService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleCreate implements ArtistHandler.
func (a *artistHandlerImpl) HandleCreate(c echo.Context) error {
	var artist req.ArtistCreateReq
	if err := c.Bind(&artist); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body"+err.Error())
	}

	if err := validate.Struct(&artist); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validation Error"+err.Error())
	}

	if err := a.createUseCase.Execute(c.Request().Context(), mapper.TransformCreateReqToEntity(artist)); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Failed to create artist"+err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", nil)
}
