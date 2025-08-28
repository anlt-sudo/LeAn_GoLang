package main

import (
	"fmt"
	"strconv"
)

func sumValidNumbers(numbers []string) {
	if len(numbers) == 0 {
		fmt.Println("Mảng không có giá trị tồn tại!")
		return
	}

	var sum int

	for _, value := range numbers {
		if num, err := strconv.Atoi(value); err != nil {
			fmt.Printf("Lỗi chuyển đổi: '%s' không phải là số!\n", value)
		} else {
			sum += num
		}
	}

	fmt.Printf("👉 Tổng các số hợp lệ là: %d\n", sum)
}

func getGrade(score int) string {
	if score < 0 || score > 100 {
		return "Điểm không hợp lệ!"
	}

	switch {
	case score >= 90:
		return "Xuất sắc"
	case score >= 80:
		return "Giỏi"
	case score >= 65:
		return "Khá"
	case score >= 50:
		return "Trung bình"
	default:
		return "Yếu"
	}
}

func main() {
	var numbers = []string{"10", "20", "hello", "30", "world", "5"}
	sumValidNumbers(numbers)

	fmt.Println(getGrade(95))  // Xuất sắc
	fmt.Println(getGrade(82))  // Giỏi
	fmt.Println(getGrade(70))  // Khá
	fmt.Println(getGrade(55))  // Trung bình
	fmt.Println(getGrade(40))  // Yếu
	fmt.Println(getGrade(-5))  // Điểm không hợp lệ!
}