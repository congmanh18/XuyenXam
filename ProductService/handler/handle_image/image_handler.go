package handle_image

import (
	"github.com/congmanh18/XuyenXam/ProductService/usecase/image_usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type ImageHandler interface {
	HandleAddImage(c echo.Context) error
	HandleDeleteImage(c echo.Context) error
	HandleGetImageByProductID(c echo.Context) error
}

type imageHandleImpl struct {
	addImageUseCase    image_usecase.AddImageUseCase
	deleteImageUseCase image_usecase.DeleteImageUseCase
	getImageUseCase    image_usecase.GetImageUseCase
}

type ImageHandleInject struct {
	AddImageUseCase    image_usecase.AddImageUseCase
	DeleteImageUseCase image_usecase.DeleteImageUseCase
	GetImageUseCase    image_usecase.GetImageUseCase
}

func NewImageHandle(inject ImageHandleInject) ImageHandler {
	return &imageHandleImpl{
		addImageUseCase:    inject.AddImageUseCase,
		deleteImageUseCase: inject.DeleteImageUseCase,
		getImageUseCase:    inject.GetImageUseCase,
	}
}
