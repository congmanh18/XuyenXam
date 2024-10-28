package req

type CategoryCreateReq struct {
	Name        *string `json:"name" validate:"required"`
	Description *string `json:"description" validate:"required"`
}

type CategoryUpdateReq struct {
	Name        *string `json:"name" validate:"required"`
	Description *string `json:"description" validate:"required"`
}
