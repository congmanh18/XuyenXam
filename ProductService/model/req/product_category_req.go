package req

type ProductCategoryReq struct {
	ProductID  *string `json:"product_id" validate:"required"`
	CategoryID *string `json:"category_id" validate:"required"`
}
