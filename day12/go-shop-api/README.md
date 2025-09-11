# Go Shop API

REST API viết bằng **Golang (Gin + GORM)** để quản lý **Danh mục (Category)** và **Sản phẩm (Product)**.  
API hỗ trợ CRUD, tìm kiếm, validation bằng DTO, và cấu trúc theo **clean architecture**.

## Mô tả

- **Category**: thêm / sửa / xóa / tìm kiếm theo tên.
- **Product**: CRUD (tương tự Category).
- **DTO**: tách struct request/response, có validation (`binding:"required"`...).
- **GORM**: kết nối MySQL/MariaDB, tự động migrate bảng.
- **Gin**: framework web để xử lý routing + middleware.
- **Air**: hỗ trợ reload code khi thay đổi (giống Spring DevTools).

## Cấu trúc thư mục

```cc
go-shop-api/
├─ cmd/
│ └─ main.go # entrypoint
├─ config/
│ └─ db.go # kết nối DB (gorm + mysql)
| |__env.example # file mẫu cấu hình env
├─ internal/
│ ├─ model/ # entity ánh xạ DB (Category, Product)
│ ├─ dto/ # struct request/response (CategoryRequest, CategoryResponse)
│ ├─ repository/ # thao tác DB (CategoryRepository, ProductRepository)
│ ├─ service/ # business logic (CategoryService, ProductService)
│ ├─ handler/ # REST handler (CategoryHandler, ProductHandler)
│ └─ router/ # khởi tạo router + inject dependency
├─ go.mod
├─ go.sum
└─ README.md
```

## 🚀 Cách chạy

### 1. Clone dự án

```bash
git clone https://github.com/your-repo/go-shop-api.git
cd go-shop-api

```

### 2. Cài dependency go mod tidy

### 3. Chạy server

go run cmd/main.go

### 4. Chạy server kèm hot reload (Air)

go install github.com/air-verse/air@latest
air

### Server chạy tại:http://localhost:8080/api/v1

### API Endpoints

#### Category

Method Endpoint Mô tả
GET /api/v1/categories Lấy tất cả danh mục
GET /api/v1/categories/1 Lấy danh mục theo ID
POST /api/v1/categories Tạo danh mục mới
PUT /api/v1/categories/1 Cập nhật danh mục
DELETE /api/v1/categories/1 Xóa danh mục
GET /api/v1/categories/search?q=laptop Tìm kiếm theo tên
Product

(Tương tự Category: CRUD + search + filter theo categoryId)

### Ví dụ cURL

Tạo Category
curl -X POST http://localhost:8080/api/v1/categories \
 -H "Content-Type: application/json" \
 -d '{"name":"Laptop","description":"Máy tính xách tay"}'

Tìm kiếm Category
curl "http://localhost:8080/api/v1/categories/search?q=laptop"

```Kết quả
Tạo thành công (201 Created)
{
"id": 1,
"name": "Laptop",
"description": "Máy tính xách tay",
"created_at": "2025-09-10T17:30:00Z",
"updated_at": "2025-09-10T17:30:00Z"
}
```

```Tìm kiếm
[
{
"id": 1,
"name": "Laptop",
"description": "Máy tính xách tay",
"created_at": "2025-09-10T17:30:00Z",
"updated_at": "2025-09-10T17:30:00Z"
}
]
```
