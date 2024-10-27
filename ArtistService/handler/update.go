package handler

import (
	"net/http"

	"github.com/congmanh18/XuyenXam/ArtistService/mapper"
	"github.com/congmanh18/XuyenXam/ArtistService/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleUpdate implements ArtistHandler.
func (a *artistHandlerImpl) HandleUpdate(c echo.Context) error {
	var artist req.ArtistUpdateReq

	if err := c.Bind(&artist); err != nil {
		return response.Error(c, http.StatusBadRequest, "Invalid request body"+err.Error())
	}

	if err := validate.Struct(&artist); err != nil {
		return response.Error(c, http.StatusBadRequest, "Validation Error"+err.Error())
	}

	if err := a.updateUseCase.Execute(c.Request().Context(), mapper.TransformUpdateReqToEntity(artist)); err != nil {
		return response.Error(c, http.StatusInternalServerError, "Internal Server Error")
	}

	return response.OK(c, http.StatusOK, "OK", nil)
}

// func (a *artistHandlerImpl) HandleUpdate(c echo.Context) error {
// 	existing := c.Get("artist").(entity.Artist)
// 	var new req.ArtistUpdateReq

// 	if err := c.Bind(&new); err != nil {
// 		return response.Error(c, http.StatusBadRequest, "Invalid request body"+err.Error())
// 	}

// 	if err := validate.Struct(&new); err != nil {
// 		return response.Error(c, http.StatusBadRequest, "Validation Error"+err.Error())
// 	}

// 	artist := mapper.MergeArtistEntity(&existing, mapper.TransformUpdateReqToEntity(new))

// 	if err := a.updateUseCase.Execute(c.Request().Context(), artist); err != nil {
// 		return response.Error(c, http.StatusInternalServerError, "Internal Server Error")
// 	}

// 	return response.OK(c, http.StatusOK, "OK", nil)
// }
