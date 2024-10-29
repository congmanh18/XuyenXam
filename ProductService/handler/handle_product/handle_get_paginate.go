package handle_product

import (
	"net/http"
	"strconv"

	"github.com/congmanh18/lucas-core/record"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// Các cải tiến trong code:
// Sử dụng c.QueryParam() để lấy giá trị từ URL query parameters
// Chuyển đổi string sang int bằng strconv.Atoi()
// Thêm validation:
// Page phải lớn hơn hoặc bằng 1
// Limit phải lớn hơn hoặc bằng 1
// Giới hạn limit tối đa là 100 để tránh quá tải hệ thống
// Bây giờ bạn có thể gọi API với các tham số như sau:
// GET /products?page=2&limit=20
// GET /products?page=1&limit=50
// GET /products (sẽ sử dụng giá trị mặc định: page=1, limit=10)
// Nếu người dùng nhập giá trị không hợp lệ như:
// page=0 hoặc page=-1 → sẽ dùng mặc định là 1
// limit=0 hoặc limit=-1 → sẽ dùng mặc định là 10
// limit=200 → sẽ giới hạn xuống còn 100
func (h *productHandleImpl) HandleGetPaginate(c echo.Context) error {
	pagination := record.Pagination{}

	// Lấy page từ query param, mặc định là 1 nếu không có hoặc không hợp lệ
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// Lấy limit từ query param, mặc định là 10 nếu không có hoặc không hợp lệ
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Giới hạn limit tối đa để tránh quá tải
	if limit > 100 {
		limit = 100
	}

	pagination.Page = page
	pagination.Limit = limit

	products, err := h.getPaginateUseCase.Execute(c.Request().Context(), &pagination)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	return response.OK(c, http.StatusOK, "OK", products)
}
