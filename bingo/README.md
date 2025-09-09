# Bài tập: Ứng dụng Bingo

## Yêu cầu

- **Tạo page bingo**: Sinh ngẫu nhiên một bảng bingo 5x5, xuất ra file `bingo_output.txt` (5 dòng đầu là bảng bingo).
- **Quay số**: Mỗi 2 giây random một số, xuất ra file (dòng 6, các số cách nhau bằng dấu cách).
- **Kiểm tra bingo**: Nếu bảng đã bingo (ngang, dọc hoặc chéo), dừng quay số.
- **Xuất kết quả**:
  - Dòng 7: Xuất ra dòng thông báo đã bingo (ngang/dọc/chéo nào).
  - Dòng 8-12: Xuất lại bảng bingo, số nào đã xuất hiện thì thay bằng 0, các ô bingo đánh dấu đặc biệt (theo code mẫu là ký tự A hoặc [..]).

## Cấu trúc dự án

```Cau truc du an
bingo/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   └── service/
│       └── service.go
└── README.md
```

## Kết quả demo

Ảnh minh họa kết quả file output:

![Kết quả bingo](./demo/ketqua.png)

## Hướng dẫn chạy

1. Build và chạy chương trình:
   ```bash
   go run ./cmd/app
   ```
2. Xem kết quả trong file `bingo_output.txt`.

---

- Mỗi lần chạy sẽ ra một bảng bingo và kết quả khác nhau.
- Có thể thay đổi logic đánh dấu hoặc xuất file theo ý muốn.
