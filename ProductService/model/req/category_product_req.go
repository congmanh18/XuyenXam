package req

type ProductCategoryCreateReq struct {
	ProductID  *string `json:"product_id" validate:"required"`
	CategoryID *string `json:"category_id" validate:"required"`
}

type ProductCategoryUpdateReq struct {
	ProductID  *string `json:"product_id" validate:"required"`
	CategoryID *string `json:"category_id" validate:"required"`
}
