# Go Generic Example: Stack

Ví dụ này minh họa cách sử dụng generic trong Go để xây dựng cấu trúc dữ liệu Stack có thể lưu trữ bất kỳ kiểu dữ liệu nào.

## Mã nguồn

File `main.go` định nghĩa một struct `Stack[T any]` với các phương thức:

- `Push(item T)`: Thêm phần tử vào stack.
- `Pop() (T, bool)`: Lấy phần tử trên cùng ra khỏi stack, trả về giá trị và trạng thái thành công.

Bạn có thể tạo stack cho bất kỳ kiểu dữ liệu nào, ví dụ:

```go
intStack := &Stack[int]{}
intStack.Push(10)
intStack.Push(20)
val, _ := intStack.Pop() // val = 20

stringStack := &Stack[string]{}
stringStack.Push("A")
stringStack.Push("B")
str, _ := stringStack.Pop() // str = "B"
```

## Ý nghĩa

Generic giúp bạn viết code tổng quát, tái sử dụng cho nhiều kiểu dữ liệu mà không cần lặp lại logic.
