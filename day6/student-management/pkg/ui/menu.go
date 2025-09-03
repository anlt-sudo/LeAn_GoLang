package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	service "github.com/anlt-sudo/student-management/internal/services"
)

var reader = bufio.NewReader(os.Stdin)

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ShowMenu(school *service.SchoolService) {
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Nhập lớp")
		fmt.Println("2. Nhập học sinh")
		fmt.Println("3. Xem tất cả lớp và học sinh")
		fmt.Println("4. Xem danh sách học sinh theo lớp")
		fmt.Println("5. Thoát")

		var choice int
		fmt.Print("Chọn chức năng: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			className := input("Nhập tên lớp: ")
			if err := school.AddClass(className); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("✅ Đã thêm lớp:", className)
			}
		case 2:
			studentName := input("Nhập tên học sinh: ")
			fmt.Println("Danh sách lớp hiện có:")
			for _, c := range school.GetAllData() {
				fmt.Printf(" - %s\n", c.Name)
			}
			className := input("Thuộc lớp nào? ")
			if err := school.AddStudent(studentName, className); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("✅ Đã thêm học sinh:", studentName)
			}
		case 3:
			fmt.Println("\n===== DANH SÁCH LỚP =====")
			for _, c := range school.GetAllData() {
				fmt.Printf("Lớp: %s (%d học sinh)\n", c.Name, c.NumStudents)
				students, _ := school.GetStudentsByClass(c.Name)
				for _, st := range students {
					fmt.Printf(" - %s\n", st.Name)
				}
				fmt.Println()
			}
		case 4:
			className := input("Nhập tên lớp cần xem: ")
			students, err := school.GetStudentsByClass(className)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("\nDanh sách học sinh lớp %s:\n", className)
				for _, st := range students {
					fmt.Printf(" - %s\n", st.Name)
				}
			}
		case 5:
			fmt.Println("👋 Thoát chương trình...")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ, vui lòng nhập lại!")
		}
	}
}
