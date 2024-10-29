package req

type ImageReq struct {
	ProductID *string `json:"product_id" validate:"required"`
	ImageURL  *string `json:"image_url" validate:"required"`
}
