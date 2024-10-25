package handler

import (
	"github.com/congmanh18/XuyenXam/ArtistService/usecase"
	"github.com/labstack/echo/v4"
)

type ArtistHandler interface {
	HandleCreate(c echo.Context) error
	HandleFindByID(c echo.Context) error
	HandleFindAll(c echo.Context) error
	HandleUpdate(c echo.Context) error
	HandleDelete(c echo.Context) error
}

type artistHandlerImpl struct {
	createUseCase   usecase.CreateUseCase
	findByIDUseCase usecase.FindByIDUseCase
	findAllUseCase  usecase.FindAllUseCase
	updateUseCase   usecase.UpdateUseCase
	deleteUseCase   usecase.DeleteUseCase
}

type ArtistHandlerInject struct {
	CreateUseCase   usecase.CreateUseCase
	FindByIDUseCase usecase.FindByIDUseCase
	FindAllUseCase  usecase.FindAllUseCase
	UpdateUseCase   usecase.UpdateUseCase
	DeleteUseCase   usecase.DeleteUseCase
}
