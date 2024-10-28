package req

type OrderCreateReq struct {
	CustomerID *string              `json:"customer_id" validate:"required"`                             // ID khách hàng, bắt buộc
	ArtistID   *string              `json:"artist_id" validate:"required"`                               // ID nghệ nhân, bắt buộc
	Status     *string              `json:"status" validate:"required,oneof=pending confirmed canceled"` // Trạng thái hợp lệ
	TotalPrice *float64             `json:"total_price" validate:"required,gt=0"`                        // Giá trị đơn hàng, phải lớn hơn 0
	Items      []OrderItemCreateReq `json:"items" validate:"required,dive"`                              // Danh sách các mục trong đơn hàng
}

type OrderItemCreateReq struct {
	OrderID     *string  `json:"order_id" validate:"required"`      //
	ProductName *string  `json:"product_name" validate:"required"`  // Tên dịch vụ/sản phẩm
	Quantity    *int     `json:"quantity" validate:"required,gt=0"` // Số lượng, phải > 0
	Price       *float64 `json:"price" validate:"required,gt=0"`    // Giá của mỗi mục, phải > 0
}
