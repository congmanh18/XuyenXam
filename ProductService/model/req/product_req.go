package req

type ProductCreateReq struct {
	Name        *string  `json:"name" validate:"required"`
	Description *string  `json:"description" validate:"required"`
	Price       *float64 `json:"price" validate:"required,gt=0"`
}

type ProductUpdateReq struct {
	Name        *string  `json:"name" validate:"required"`
	Description *string  `json:"description" validate:"required"`
	Price       *float64 `json:"price" validate:"required,gt=0"`
}
