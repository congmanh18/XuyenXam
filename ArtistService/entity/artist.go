package entity

import (
	"github.com/congmanh18/lucas-core/record"
)

type Artist struct {
	record.BaseEntity
	Firstname    *string   `bson:"first_name"`
	Lastname     *string   `bson:"last_name"`
	Username     *string   `bson:"username"`
	Email        *string   `bson:"email"`
	PhoneNumber  *string   `bson:"phone_number"`
	Biography    *string   `bson:"biography"`    // Giới thiệu ngắn về nghệ nhân
	Styles       *[]string `bson:"styles"`       // Các phong cách xăm mà nghệ nhân chuyên
	Rating       *float64  `bson:"rating"`       // Điểm đánh giá trung bình từ khách hàng
	Availability *[]Slot   `bson:"availability"` // Lịch khả dụng
	Gallery      *[]Image  `bson:"gallery"`      // Các mẫu hình xăm đã thực hiện
	Location     *Address  `bson:"location"`     // Địa chỉ studio của nghệ nhân
	SocialLinks  *[]string `bson:"social_links"` // Liên kết mạng xã hội (Instagram, Facebook, etc.)
}

type Slot struct {
	DayOfWeek string `bson:"day_of_week"` // Ví dụ: Monday, Tuesday
	StartTime string `bson:"start_time"`  // Thời gian bắt đầu (HH:mm)
	EndTime   string `bson:"end_time"`    // Thời gian kết thúc (HH:mm)
}
type Image struct {
	URL         string `bson:"url"`         // Đường dẫn đến ảnh
	Description string `bson:"description"` // Mô tả ảnh
	UploadedAt  int64  `bson:"uploaded_at"` // Thời gian tải lên
}

type Address struct {
	Street     string `json:"street" bson:"street"`           // Địa chỉ chi tiết
	City       string `json:"city" bson:"city"`               // Thành phố
	State      string `json:"state" bson:"state"`             // Bang hoặc tỉnh
	PostalCode string `json:"postal_code" bson:"postal_code"` // Mã bưu chính
	Country    string `json:"country" bson:"country"`         // Quốc gia
}
