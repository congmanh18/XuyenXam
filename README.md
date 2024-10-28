### Các bước cấu hình Nginx trên 1 máy VPS
docker run --name postgres-17 --rm -e POSTGRES_USER=manh -e POSTGRES_PASSWORD=123456 -p 5432:5432 -it -d postgres:17

### 1. **Tổng Quan**  
---

### 2. **Phân tích Tính Năng và Lợi Ích**  

#### **Tính năng chính và Lợi ích:**
1. **Đặt lịch hẹn với nghệ nhân:**  
   - Người dùng xem thời gian trống của nghệ nhân theo lịch trực tuyến.  
   - Giảm thiểu rủi ro trùng lịch, tối ưu hóa thời gian của nghệ nhân.  

2. **Quản lý đơn order:**  
   - Người dùng gửi yêu cầu tùy chỉnh, ví dụ: thiết kế mẫu, giá cả.  
   - Nghệ nhân theo dõi và xác nhận đơn hàng kịp thời.  

3. **Đánh giá và review:**  
   - Hệ thống đánh giá chất lượng nghệ nhân qua feedback.  
   - Nâng cao uy tín nghệ nhân thông qua phản hồi tích cực, tạo niềm tin cho khách hàng mới.  

4. **Thanh toán online và thông báo nhắc lịch:**  
   - Hỗ trợ thanh toán qua nhiều phương thức như thẻ ngân hàng, ví điện tử.  
   - Nhắc nhở lịch hẹn tự động qua email/SMS giúp hạn chế lỡ hẹn.  

---

### 3. **Phân tích Microservices**

1. **Artist Service:**  
   - **Chức năng:** Quản lý thông tin nghệ nhân, hồ sơ cá nhân, lịch làm việc và các mẫu hình nghệ thuật.  
   - **Dữ liệu:** Lưu thông tin nghệ nhân, thời gian khả dụng, và các hình ảnh mẫu.  
   - **API chính:** 
     - GET /artists: Lấy danh sách nghệ nhân.
     - GET /artists/{id}/availability: Kiểm tra thời gian trống.
     - POST /artists: Thêm nghệ nhân mới.  

2. **Order Service:**  
   - **Chức năng:** Xử lý đơn đặt hàng và yêu cầu dịch vụ xăm từ người dùng.  
   - **Dữ liệu:** Thông tin đơn hàng, yêu cầu của khách, trạng thái đơn.  
   - **API chính:** 
     - POST /orders: Tạo đơn hàng mới.
     - GET /orders/{id}: Kiểm tra chi tiết đơn hàng.  

3. **Payment Service:**  
   - **Chức năng:** Thanh toán trực tuyến và xử lý hóa đơn.  
   - **Tích hợp:** Cổng thanh toán như Stripe, PayPal.  
   - **API chính:** 
     - POST /payments: Thực hiện thanh toán.
     - GET /payments/{id}: Kiểm tra trạng thái giao dịch.  

4. **Review Service:**  
   - **Chức năng:** Lưu trữ và quản lý đánh giá từ khách hàng về dịch vụ và nghệ nhân.  
   - **API chính:** 
     - POST /reviews: Thêm đánh giá mới.
     - GET /reviews/{artistId}: Lấy đánh giá của một nghệ nhân.  

---

### 4. **Kiến Trúc Công Nghệ**

1. **Backend:**  
   - **Golang:** Được chọn vì hiệu năng cao, phù hợp với microservices.  
   - **MongoDB:** Xử lý dữ liệu phi cấu trúc như hình ảnh mẫu và thông tin nghệ nhân.  

2. **Authentication:**  
   - **JWT (JSON Web Token):** Quản lý phiên đăng nhập an toàn và phân quyền (user, admin, artist).  

3. **Quản lý thông báo:**  
   - **RabbitMQ hoặc Kafka:** Quản lý thông báo theo thời gian thực, ví dụ: nhắc nhở lịch hẹn hoặc thay đổi trạng thái đơn hàng.  

4. **Containerization:**  
   - **Docker:** Đảm bảo môi trường chạy đồng nhất cho các microservices, giúp dễ triển khai và bảo trì.  

---

### 5. **Luồng Hoạt Động của Ứng Dụng**

1. **Người dùng tạo tài khoản và đăng nhập** (sử dụng JWT).  
2. **Người dùng tìm kiếm nghệ nhân** theo phong cách, vị trí, hoặc rating.  
3. **Kiểm tra thời gian trống** và **đặt lịch hẹn**.  
4. **Thanh toán online** và nhận thông báo xác nhận lịch.  
5. Nghệ nhân và người dùng **nhận thông báo nhắc nhở** trước buổi hẹn.  
6. **Sau buổi hẹn, khách hàng đánh giá nghệ nhân** và dịch vụ trên hệ thống.  

---

### 6. **Tính Năng Bổ Sung (Optional)**
- **Tích hợp Chat trực tiếp:** Giữa khách và nghệ nhân để thảo luận yêu cầu thiết kế.  
- **Địa chỉ và hướng dẫn đến studio:** Tích hợp bản đồ giúp khách tìm kiếm dễ dàng.  
- **Lịch sử dịch vụ:** Lưu lại lịch sử xăm để khách theo dõi và dễ dàng lên kế hoạch cho lần tiếp theo.  

---

### 7. **Các Yếu Tố Thành Công**

- **Trải nghiệm người dùng mượt mà:** Tìm kiếm nhanh, thanh toán tiện lợi, và thông báo nhắc nhở đầy đủ.
- **Tính tin cậy và bảo mật:** Bảo mật thanh toán và thông tin người dùng bằng JWT và giao thức HTTPS.
- **Hiệu năng tốt:** Golang và RabbitMQ/Kafka giúp ứng dụng xử lý nhiều tác vụ đồng thời mà không gặp vấn đề về hiệu suất.

---

### 8. **Khó Khăn Tiềm Ẩn và Giải Pháp**

1. **Xử lý xung đột lịch hẹn:**  
   - Sử dụng RabbitMQ để cập nhật trạng thái thời gian trống theo thời gian thực.  
2. **Thanh toán bị gián đoạn:**  
   - Tích hợp cơ chế thanh toán dự phòng, lưu trạng thái đơn hàng nếu lỗi xảy ra.  
3. **Hệ thống đánh giá sai lệch:**  
   - Sử dụng thuật toán phát hiện đánh giá spam hoặc không trung thực.

---

### 9. **Tổng Kết**
Dự án ứng dụng này mang đến **trải nghiệm thuận tiện và tin cậy** cho cả nghệ nhân lẫn khách hàng, đáp ứng nhu cầu về quản lý lịch hẹn và dịch vụ xăm hình hiệu quả. Việc áp dụng **microservices và các công nghệ tiên tiến** như Golang, MongoDB, RabbitMQ/Kafka sẽ giúp ứng dụng dễ dàng mở rộng và duy trì hiệu năng cao trong dài hạn.


```

// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs


Table artist {
  id string [primary key]
  name string
  specialization string
  orders string []
  created_at timestamp
  updated_at timestamp
}

Table customer {
  id string [primary key]
  name string
  email string
  phone_number string
  created_at timestamp
  updated_at timestamp
}

Table order {
  id string [primary key]
  artist_id string
  customer_id string
  items array
  payment string
  created_at timestamp
  updated_at timestamp
}

Table order_item {
  id integer [primary key]
  product_name string
  quantity int
  price float
  created_at timestamp
}

Table payment {
  id string [primary key]

  created_at timestamp
}

Ref: order.customer_id - customer.id // many-to-one
Ref: order.artist_id - artist.id
Ref: order_item.id > order.items
Ref: order.payment - payment.id


Table product {
  id string [primary key]
  name string
  description string
  price float
  created_at timestamp
  updated_at timestamp
}


Table category {
  id string [primary key]
  name string
  description string
}

Table product_category {
  product_id string [primary key]
  category_id string [primary key]
}

Table inventory {
  product_id string [primary key]
  quantity int
  updated_at timestamp
}

Table image {
  id string [primary key]
  product_id string 
  url string
  created_at timestamp
}


Ref: product.id - product_category.product_id // many-to-one
Ref: category.id - product_category.category_id
Ref: product.id < inventory.product_id
Ref: product.id < image.product_id


```