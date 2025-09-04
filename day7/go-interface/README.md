# Hệ thống Thông báo Đa kênh bằng Golang

Đây là một dự án Go đơn giản để minh họa các khái niệm cốt lõi về **Interface**, cách tổ chức code theo một cấu trúc thư mục hợp lý, và các nguyên tắc thiết kế phần mềm tốt như **tách biệt (decoupling)**.

Dự án xây dựng một hệ thống có thể gửi thông báo qua nhiều kênh khác nhau (Email, SMS, Slack) bằng cách sử dụng một hàm xử lý chung.

## Các khái niệm chính được minh họa

- **Interface:** Định nghĩa một "hợp đồng" chung (`Notifier`) cho các hành vi gửi thông báo.
- **Thỏa mãn Interface ngầm định:** Các `struct` (`EmailNotifier`, `SMSNotifier`,...) tự động thỏa mãn interface mà không cần khai báo tường minh.
- **Cấu trúc dự án:** Phân chia code thành các package có trách nhiệm rõ ràng (`cmd`, `internal/notification`, `internal/notifiers`).
- **`interface{}` và `Type Switch`:** Xử lý một tập hợp các đối tượng có kiểu dữ liệu khác nhau một cách an toàn.

## Cấu trúc Thư mục
code
Code

- **`cmd/app/main.go`**: Điểm khởi đầu của ứng dụng. Chịu trách nhiệm khởi tạo và kết nối các thành phần lại với nhau.
- **`internal/notification/service.go`**: Định nghĩa interface `Notifier` cốt lõi và các logic xử lý chung.
- **`internal/notifiers/`**: Chứa các triển khai cụ thể cho từng kênh thông báo (Email, SMS, Slack).

## Yêu cầu

- Go phiên bản 1.18 trở lên.

## Hướng dẫn Cài đặt và Chạy

1.  **Clone repository (hoặc tạo lại dự án):**
    Nếu bạn đang xem trên GitHub, hãy clone repository. Nếu không, hãy tạo lại cấu trúc thư mục và các file như đã mô tả ở trên.

2.  **Khởi tạo module (nếu tạo lại dự án):**
    Mở terminal ở thư mục gốc của dự án (`go-interface/`) và chạy lệnh:

    ```bash
    go mod init go-interface
    ```

3.  **Chạy ứng dụng:**
    Từ thư mục gốc, chạy lệnh sau để biên dịch và thực thi chương trình:
    ```bash
    go run ./cmd/app/main.go
    ```

## Kết quả mong muốn

Sau khi chạy, bạn sẽ thấy kết quả đầu ra trên console như sau:
--- Sending notifications via Notifier slice ---
Sending email to test@example.com: Your order has been shipped!
Sending SMS to 123-456-7890: Your order has been shipped!
Sending Slack message to #general: Your order has been shipped!
--- Processing mixed notifications ---
Sending email to test@example.com: System will be down for maintenance.
Skipping notification for number: 123
Sending SMS to 123-456-7890: System will be down for maintenance.
Unsupported notifier type: main.InvalidNotifier
Sending Slack message to #general: System will be down for maintenance.
