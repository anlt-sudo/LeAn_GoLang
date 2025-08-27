Chào bạn, dưới đây là tổng hợp các thông tin về package, import, Go Modules và cấu trúc thư mục tiêu chuẩn trong một dự án Go.

### 1. Package và Import

Trong Go, code được tổ chức thành các **package**. Mỗi tệp mã nguồn Go đều phải thuộc về một package, được khai báo ở đầu tệp bằng từ khóa `package`.

*   **Package `main`**: Đây là một package đặc biệt. Nó định nghĩa một chương trình có thể thực thi được, và hàm `main()` bên trong package này sẽ là điểm khởi đầu của chương trình.
*   **Tổ chức code**: Packages giúp bạn module hóa mã nguồn, làm cho nó dễ đọc, dễ bảo trì và tái sử dụng hơn.
*   **Export (Xuất)**: Các định danh (hàm, biến, hằng số, kiểu dữ liệu) được viết hoa chữ cái đầu sẽ được "xuất" ra ngoài. Điều này có nghĩa là các package khác có thể import và sử dụng chúng. Các định danh viết thường sẽ không được xuất và chỉ có thể sử dụng bên trong cùng một package.

**Import**

Để sử dụng code từ một package khác, bạn cần "import" nó bằng từ khóa `import`.

*   **Cú pháp**: Bạn có thể import từng package riêng lẻ hoặc import một nhóm các package.
*   **Đường dẫn import**: Chuỗi ký tự trong câu lệnh import được gọi là đường dẫn import. Nó xác định package nào sẽ được import. Đối với các package trong thư viện chuẩn của Go, bạn chỉ cần dùng tên ngắn gọn như `"fmt"` hay `"math/rand"`. Đối với các package của bên thứ ba, đường dẫn này thường là URL của kho mã nguồn, ví dụ: `"github.com/sirupsen/logrus"`.
*   **Sử dụng**: Sau khi import, tên package sẽ được dùng để truy cập các thành phần đã được export từ package đó. Ví dụ, sau khi `import "fmt"`, bạn có thể gọi hàm `Println` bằng cách viết `fmt.Println()`.

### 2. Go Modules

Go Modules là hệ thống quản lý dependency (các thư viện phụ thuộc) chính thức của Go, được giới thiệu từ phiên bản Go 1.11. Một module về cơ bản là một tập hợp các package Go được phiên bản hóa cùng nhau. Go Modules cho phép bạn đặt dự án ở bất kỳ đâu trên máy tính mà không cần phụ thuộc vào `GOPATH` như trước đây.

File `go.mod` là trung tâm của một module. Nó định nghĩa đường dẫn của module và liệt kê tất cả các dependency cần thiết.

#### **go mod init**

Lệnh này dùng để khởi tạo một module mới.

*   **Chức năng**: `go mod init` tạo ra một tệp `go.mod` trong thư mục hiện tại, đánh dấu đây là thư mục gốc của một module.
*   **Cách dùng**: Bạn cần cung cấp một đường dẫn module, thường là URL của kho chứa mã nguồn của bạn.
    ```bash
    go mod init github.com/ten-cua-ban/ten-du-an
    ```
    Lệnh này sẽ tạo ra một tệp `go.mod` với nội dung tương tự:
    ```
    module github.com/ten-cua-ban/ten-du-an

    go 1.17
    ```

#### **go get**

Lệnh `go get` được dùng để quản lý các dependency.

*   **Chức năng**: Nó có thể thêm, cập nhật hoặc xóa các dependency trong tệp `go.mod` của bạn. Khi bạn `go get` một package, Go sẽ tự động tải mã nguồn của package đó về máy.
*   **Cách dùng**:
    *   Để thêm một dependency mới hoặc cập nhật lên phiên bản mới nhất:
        ```bash
        go get github.com/gorilla/mux
        ```
    *   Để tải một phiên bản cụ thể:
        ```bash
        go get github.com/gorilla/mux@v1.8.0
        ```
    *   **Lưu ý**: Kể từ các phiên bản Go gần đây, vai trò của `go get` đã thay đổi. Nó chủ yếu dùng để điều chỉnh các dependency trong `go.mod`. Để cài đặt các tệp thực thi, lệnh `go install` được khuyến khích sử dụng.

#### **go mod tidy**

Lệnh này dùng để "dọn dẹp" tệp `go.mod` và `go.sum`.

*   **Chức năng**: `go mod tidy` đảm bảo rằng tệp `go.mod` khớp với mã nguồn trong dự án của bạn. Nó sẽ:
    *   Thêm các module còn thiếu mà code của bạn cần để build.
    *   Loại bỏ các module không được sử dụng.
    *   Cập nhật tệp `go.sum`, nơi chứa checksum của các dependency để đảm bảo tính toàn vẹn.
*   **Cách dùng**:
    ```bash
    go mod tidy
    ```

### 3. Cấu trúc thư mục của một dự án Go tiêu chuẩn

Không có một cấu trúc dự án "chính thức" duy nhất do đội ngũ Go định nghĩa. Cấu trúc có thể thay đổi tùy thuộc vào quy mô và loại dự án. Tuy nhiên, cộng đồng đã hình thành một số quy ước và mẫu cấu trúc chung.

Đối với một dự án nhỏ hoặc khi mới bắt đầu, bạn chỉ cần một tệp `main.go` và `go.mod` là đủ. Khi dự án lớn dần, bạn có thể áp dụng cấu trúc sau:

**/cmd**
Thư mục này chứa mã nguồn cho các ứng dụng chính (các tệp thực thi) của bạn. Tên của mỗi thư mục con bên trong `/cmd` thường trùng với tên của tệp thực thi mà bạn muốn tạo.

**/internal**
Đây là nơi chứa các package và code riêng tư của ứng dụng. Mã nguồn trong thư mục `internal` không thể được import bởi các dự án bên ngoài. Đây là một quy tắc được chính công cụ của Go thực thi.

**/pkg**
Thư mục này chứa các thư viện công khai, có thể được sử dụng bởi các dự án khác. Hãy cân nhắc kỹ trước khi đặt code vào đây, vì nó tạo ra một cam kết không chính thức về việc duy trì tính ổn định cho người dùng bên ngoài.

**/api**
Chứa các tệp định nghĩa hợp đồng API như tệp tin OpenAPI/Swagger, định nghĩa Protocol Buffer.

**/web**
Chứa các tài sản dành riêng cho ứng dụng web như tệp tĩnh, template.

**/configs**
Chứa các tệp cấu hình mẫu hoặc mặc định.

**/scripts**
Chứa các script để thực hiện các tác vụ khác nhau như build, cài đặt, phân tích, v.v.

**Ví dụ về cấu trúc thư mục:**
```
my-project/
├── go.mod
├── go.sum
├── cmd/
│   └── my-app/
│       └── main.go
├── internal/
│   ├── auth/
│   │   └── auth.go
│   └── user/
│       └── user.go
├── pkg/
│   └── database/
│       └── database.go
├── api/
│   └── swagger.yaml
├── web/
│   ├── static/
│   └── templates/
└── configs/
    └── config.yaml
```
