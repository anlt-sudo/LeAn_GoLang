# Go Logging Demo

> Ví dụ sử dụng hai thư viện logging phổ biến trong Go: **zerolog** và **zap**.

## Mục tiêu

- So sánh cách sử dụng và kết quả log giữa zerolog và zap.
- Demo log các level khác nhau (debug, info, warn, error).
- Gọi trực tiếp handler để log thông tin xử lý request mà không cần chạy server.


## Kết quả log

### 1. Log với zerolog

![Log với zerolog](./demo/zerolog-demo.png)

### 2. Log với zap

![Log với zap](./demo/zap-demo.png)

> Ảnh minh họa: hai file ảnh kết quả log thực tế (pi1, pi2) đã được đính kèm trong repo.

## Ghi chú

- **zerolog**: log nhanh, cấu hình đơn giản, output mặc định là JSON, có thể chuyển sang dạng console cho dễ đọc.
- **zap**: log mạnh mẽ, nhiều tính năng nâng cao, hỗ trợ cả production và development mode, output mặc định là JSON.
- Handler trong ví dụ vẫn dùng zerolog để minh họa việc tích hợp nhiều logger trong cùng một dự án.

## Tham khảo

- https://github.com/rs/zerolog
- https://github.com/uber-go/zap
