package req

type InventoryReq struct {
	Quantity *int `json:"quantity" validate:"required,gt=0"`
}
