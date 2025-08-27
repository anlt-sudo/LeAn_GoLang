package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type UserManager struct {
	users  map[int]User
	nextID int
}

func NewUserManager() *UserManager {
	return &UserManager{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (um *UserManager) AddUser(name, email string, age int) User {
	newUser := User{
		ID:    um.nextID,
		Name:  name,
		Email: email,
		Age:   age,
	}
	um.users[um.nextID] = newUser
	um.nextID++
	return newUser
}


func (um *UserManager) GetUser(id int) (User, bool) {
	user, ok := um.users[id]
	return user, ok
}


func (um *UserManager) UpdateUser(id, age int, name, email string) bool {
	if _, ok := um.users[id]; !ok {
		return false
	}
	updatedUser := User{ID: id, Name: name, Email: email, Age: age}
	um.users[id] = updatedUser
	return true
}

func (um *UserManager) DeleteUser(id int) bool {
	if _, ok := um.users[id]; !ok {
		return false
	}
	delete(um.users, id)
	return true
}

func (um *UserManager) ListUsers() []User {
	userList := make([]User, 0, len(um.users))
	for _, user := range um.users {
		userList = append(userList, user)
	}
	return userList
}


func main() {
	manager := NewUserManager()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--- Chào mừng đến với Chương trình Quản lý Người dùng ---")

	for {
		fmt.Println("\nVui lòng chọn một chức năng:")
		fmt.Println("1. Thêm người dùng mới")
		fmt.Println("2. Liệt kê tất cả người dùng")
		fmt.Println("3. Tìm người dùng theo ID")
		fmt.Println("4. Cập nhật thông tin người dùng")
		fmt.Println("5. Xóa người dùng")
		fmt.Println("6. Thoát")
		fmt.Print("Lựa chọn của bạn: ")

		choiceStr, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceStr)

		switch choice {
		case "1": // Thêm người dùng
			fmt.Print("Nhập tên: ")
			name, _ := reader.ReadString('\n')

			fmt.Print("Nhập email: ")
			email, _ := reader.ReadString('\n')

			fmt.Print("Nhập tuổi: ")
			ageStr, _ := reader.ReadString('\n')
			age, err := strconv.Atoi(strings.TrimSpace(ageStr))
			if err != nil {
				fmt.Println("Lỗi: Tuổi không hợp lệ.")
				continue
			}

			user := manager.AddUser(strings.TrimSpace(name), strings.TrimSpace(email), age)
			fmt.Printf("=> Đã thêm người dùng thành công với ID: %d\n", user.ID)

		case "2": // Liệt kê người dùng
			users := manager.ListUsers()
			if len(users) == 0 {
				fmt.Println("=> Hiện không có người dùng nào.")
				continue
			}
			fmt.Println("--- Danh sách người dùng ---")
			for _, u := range users {
				fmt.Printf("ID: %d, Tên: %s, Email: %s, Tuổi: %d\n", u.ID, u.Name, u.Email, u.Age)
			}

		case "3": // Tìm người dùng
			fmt.Print("Nhập ID người dùng cần tìm: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Lỗi: ID không hợp lệ.")
				continue
			}

			user, ok := manager.GetUser(id)
			if ok {
				fmt.Println("=> Tìm thấy người dùng:")
				fmt.Printf("ID: %d, Tên: %s, Email: %s, Tuổi: %d\n", user.ID, user.Name, user.Email, user.Age)
			} else {
				fmt.Printf("=> Không tìm thấy người dùng với ID: %d\n", id)
			}

		case "4": // Cập nhật người dùng
			fmt.Print("Nhập ID người dùng cần cập nhật: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Lỗi: ID không hợp lệ.")
				continue
			}

			// Kiểm tra user có tồn tại không
			if _, ok := manager.GetUser(id); !ok {
				fmt.Printf("=> Không tìm thấy người dùng với ID: %d để cập nhật.\n", id)
				continue
			}

			fmt.Print("Nhập tên mới: ")
			name, _ := reader.ReadString('\n')
			fmt.Print("Nhập email mới: ")
			email, _ := reader.ReadString('\n')
			fmt.Print("Nhập tuổi mới: ")
			ageStr, _ := reader.ReadString('\n')
			age, err := strconv.Atoi(strings.TrimSpace(ageStr))
			if err != nil {
				fmt.Println("Lỗi: Tuổi không hợp lệ.")
				continue
			}

			if manager.UpdateUser(id, age, strings.TrimSpace(name), strings.TrimSpace(email)) {
				fmt.Println("=> Cập nhật thông tin người dùng thành công.")
			}

		case "5": // Xóa người dùng
			fmt.Print("Nhập ID người dùng cần xóa: ")
			idStr, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Lỗi: ID không hợp lệ.")
				continue
			}

			if manager.DeleteUser(id) {
				fmt.Println("=> Đã xóa người dùng thành công.")
			} else {
				fmt.Printf("=> Không tìm thấy người dùng với ID: %d để xóa.\n", id)
			}

		case "6": // Thoát
			fmt.Println("Tạm biệt!")
			os.Exit(0)

		default:
			fmt.Println("Lựa chọn không hợp lệ, vui lòng thử lại.")
		}
	}
}