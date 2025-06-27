# 🧪 REST API: Blog Post Management 

## 📌 Mục tiêu
Xây dựng một RESTful API bằng Golang để quản lý bài viết blog. Bài test này giúp kiểm tra kiến thức của bạn về:
- HTTP methods
- JSON request/response
- Golang struct, router, middleware
- GORM & MySQL
- Clean code & code structure

---

## 🚀 Yêu cầu chức năng

1. **Tạo bài viết mới**  
   `POST /posts`  
   Request body:
   ```json
   {
     "title": "Tiêu đề",
     "content": "Nội dung bài viết",
     "author": "Tên tác giả"
   }
   ```

2. **Lấy danh sách tất cả bài viết**  
   `GET /posts`
   Request Query:    
   ```params
   {
     "page": "1",
     "limit": "10",
   }
   ```
   Response body:
   ```json
   {
     status: true,
     data: [],
     total: 0,
   }
   ```

3. **Lấy chi tiết 1 bài viết theo ID**  
   `GET /posts/:id`

4. **Cập nhật bài viết**  
   `PUT /posts/:id`  
   Request body (giống như tạo mới)

5. **Xoá bài viết**  
   `DELETE /posts/:id`

6. **Xoá nhiều bài viết theo id**  
   `DELETE /posts/many/:id`
---

## 🧱 Database structure

Bảng `posts` với các trường:

| Field      | Type         |
|------------|--------------|
| id         | uint, PK     |
| title      | varchar(255) |
| content    | text         |
| author     | varchar(100) |
| created_at | datetime     |

> Sử dụng GORM để migrate tự động hoặc viết SQL tay đều được.

---

## ⚙️ Công nghệ yêu cầu

| Tech         | Yêu cầu                      |
|--------------|------------------------------|
| Golang       | >= 1.18                      |
| Router       | `gin`, `net/http`, fiber     |
| ORM          | `gorm.io/gorm` + `MySQL`     |
| Structuring  | Theo pattern `MVC/service`   |

---

## 📂 Gợi ý cấu trúc thư mục

```
.
├── cmd/                 # main.go
├── config/              # config env/db
├── controllers/         # post_controller.go
├── models/              # post.go
├── repositories/        # post_repository.go
├── routes/              # routes.go
├── services/            # post_service.go
├── utils/               # helpers, response wrapper
├── go.mod
└── README.md
```

---

## 📦 Cách chạy project

```bash
# 1. Clone repo
git clone <link-repo>

# 2. Khởi tạo file .env
cp .env.example .env

# 3. Cài đặt thư viện
go mod tidy

# 4. Chạy app
go run cmd/main.go
```

---

## ✅ Kiểm tra

- Test API bằng Postman hoặc curl
- Đảm bảo status code chuẩn (200, 201, 400, 404...)
- Validate dữ liệu (không cho title trống,...)
- Gợi ý: tạo `swagger.yaml` nếu biết

---

## ⭐ Điểm cộng nếu có

- Pagination `GET /posts?page=1&limit=10`
- JWT Auth cơ bản
- Swagger Docs
- Unit test cho controller/service

---

## 🧠 Lưu ý đánh giá

| Tiêu chí | Có tính điểm |
|----------|--------------|
| Đúng yêu cầu chức năng | ✅ |
| Cấu trúc project rõ ràng | ✅ |
| Xử lý lỗi tốt | ✅ |
| Format JSON đẹp, đúng chuẩn | ✅ |
| Commit rõ ràng, gọn gàng | ✅ |

---
