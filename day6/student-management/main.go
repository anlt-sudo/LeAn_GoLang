package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Class struct {
	Name        string
	NumStudents int
}

type Student struct {
	Name      string
	ClassName string
}

var classes []*Class
var students []*Student

var classSet = make(map[string]struct{})

var reader = bufio.NewReader(os.Stdin)

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func addClass() {
	className := input("Nhập tên lớp: ")

	key := strings.ToLower(className)

	if _, exists := classSet[key]; exists {
		fmt.Println("❌ Lỗi: Lớp đã tồn tại")
		return
	}

	class := &Class{
		Name:        className,
		NumStudents: 0,
	}
	classes = append(classes, class)

	classSet[key] = struct{}{}

	fmt.Println("✅ Đã thêm lớp:", className)
}

func updateClassStudentCount(className string) {
	for _, c := range classes {
		if strings.EqualFold(className, c.Name) {
			count := 0
			for _, s := range students {
				if strings.EqualFold(s.ClassName, c.Name) {
					count++
				}
			}
			c.NumStudents = count
			return
		}
	}
}

func addStudent() {
	if len(classes) == 0 {
		fmt.Println("❌ Chưa có lớp nào, vui lòng nhập lớp trước!")
		return
	}

	studentName := input("Nhập tên học sinh: ")

	fmt.Println("Danh sách lớp hiện có:")
	for index, class := range classes {
		fmt.Printf("%d. %s\t", index+1, class.Name)
	}
	className := input("\nThuộc lớp nào? ")

	if _, exists := classSet[strings.ToLower(className)]; !exists {
		fmt.Println("❌ Lỗi: Tên lớp không tồn tại trong danh sách")
		return
	}

	student := &Student{
		Name:      studentName,
		ClassName: className,
	}
	students = append(students, student)

	updateClassStudentCount(className)

	fmt.Println("✅ Đã thêm học sinh:", studentName)
}

func showData() {
	fmt.Println("\n===== DANH SÁCH LỚP =====")
	for _, c := range classes {
		updateClassStudentCount(c.Name)
		fmt.Printf("Lớp: %s (%d học sinh)\n", c.Name, c.NumStudents)
		for _, s := range students {
			if strings.EqualFold(s.ClassName, c.Name) {
				fmt.Printf(" - %s\n", s.Name)
			}
		}
		fmt.Println()
	}
}

func showDataByClassName(className string) {
	if _, exists := classSet[strings.ToLower(className)]; !exists {
		fmt.Println("❌ Không tìm thấy lớp", className)
		return
	}

	for _, c := range classes {
		if strings.EqualFold(className, c.Name) {
			updateClassStudentCount(c.Name)
			fmt.Printf("\nLớp: %s (%d học sinh)\n", c.Name, c.NumStudents)
			for _, s := range students {
				if strings.EqualFold(s.ClassName, c.Name) {
					fmt.Printf(" - %s\n", s.Name)
				}
			}
			fmt.Println()
			return
		}
	}
}

func main() {
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
			addClass()
		case 2:
			addStudent()
		case 3:
			showData()
		case 4:
			className := input("Nhập tên lớp cần xem: ")
			showDataByClassName(className)
		case 5:
			fmt.Println("👋 Thoát chương trình...")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ, vui lòng nhập lại!")
		}
	}
}
