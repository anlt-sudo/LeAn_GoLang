package main

import (
	"fmt"
)

// Bước 1: Định nghĩa kiểu cho các bước xử lý
// Ứng dụng: Hàm là giá trị (Functions as Values)
type operation func(int) (int, error)

// Bước 2: Xây dựng hàm processPipeline
// Ứng dụng: Hàm Variadic, Multiple Returns, Defer
func processPipeline(data []int, ops ...operation) ([]int, error) {
	fmt.Println("Pipeline processing started...")
	// Ứng dụng: Defer để đảm bảo thông báo này luôn được in ra cuối cùng.
	defer fmt.Println("Pipeline processing finished.")

	processedData := make([]int, len(data))

	for i, value := range data {
		currentValue := value
		for _, op := range ops {
			var err error
			currentValue, err = op(currentValue)
			// Ứng dụng: Xử lý lỗi từ multiple returns
			if err != nil {
				// Dừng ngay lập tức và trả về lỗi
				return nil, fmt.Errorf("lỗi tại giá trị %d: %v", value, err)
			}
		}
		processedData[i] = currentValue
	}

	return processedData, nil
}

// Bước 3: Tạo các hàm xử lý
func double(n int) (int, error) {
	return n * 2, nil
}

func addFive(n int) (int, error) {
	return n + 5, nil
}

func rejectEvens(n int) (int, error) {
	if n%2 == 0 {
		return 0, fmt.Errorf("giá trị %d là số chẵn, đã bị từ chối", n)
	}
	return n, nil
}

// Bước 4: Tạo một "Factory" cho các hàm xử lý
// Ứng dụng: Closure
func createMultiplier(factor int) operation {
	// Hàm ẩn danh bên dưới là một closure.
	// Nó "nhớ" giá trị của 'factor' từ môi trường mà nó được tạo ra.
	return func(n int) (int, error) {
		return n * factor, nil
	}
}

// Bước 5: Viết hàm main để kiểm thử
func main() {
	initialData := []int{1, 3, 5, 7}
	fmt.Printf("Dữ liệu ban đầu: %v\n\n", initialData)

	// --- Kịch bản 1: Thành công ---
	fmt.Println("--- Kịch bản 1: Thành công ---")
	multiplyBy3 := createMultiplier(3) // Tạo một hàm xử lý từ closure
	result1, err1 := processPipeline(initialData, double, addFive, multiplyBy3)
	if err1 != nil {
		fmt.Println("Lỗi:", err1)
	} else {
		fmt.Printf("Kết quả: %v\n", result1)
	}
	// Giải thích: (1*2+5)*3=21, (3*2+5)*3=33, ...

	fmt.Println("\n----------------------------------\n")

	// --- Kịch bản 2: Thất bại ---
	fmt.Println("--- Kịch bản 2: Thất bại ---")
	dataWithEven := []int{1, 2, 3}
	fmt.Printf("Dữ liệu ban đầu: %v\n", dataWithEven)
	_, err2 := processPipeline(dataWithEven, addFive, rejectEvens)
	if err2 != nil {
		fmt.Println("Lỗi:", err2)
	}
	// Giải thích: 1+5=6. rejectEvens(6) sẽ gây lỗi.

	fmt.Println("\n----------------------------------\n")

	// --- Kịch bản 3: Với hàm ẩn danh ---
	fmt.Println("--- Kịch bản 3: Với hàm ẩn danh ---")
	dataForCapping := []int{10, 50, 80}
	fmt.Printf("Dữ liệu ban đầu: %v\n", dataForCapping)
	// Ứng dụng: Hàm Ẩn danh (Anonymous Function)
	capAt100 := func(n int) (int, error) {
		if n > 100 {
			return 100, nil
		}
		return n, nil
	}
	result3, err3 := processPipeline(dataForCapping, double, capAt100)
	if err3 != nil {
		fmt.Println("Lỗi:", err3)
	} else {
		fmt.Printf("Kết quả: %v\n", result3)
	}
}