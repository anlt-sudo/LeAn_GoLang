package main

import "fmt"

type User struct {
	Name     string
	Age      int
	IsActive bool
}


func celebrateBirthday(u *User) {
	fmt.Printf("\nChúc mừng sinh nhật %s!\n", u.Name)
	u.Age++
}


func updateStatus(u *User, newStatus bool) {
	u.IsActive = newStatus
}


func main() {

	user := User{
		Name:     "Alice",
		Age:      30,
		IsActive: true,
	}

	fmt.Println("--- Trạng thái ban đầu ---")
	fmt.Printf("Tên: %s, Tuổi: %d, Hoạt động: %t\n", user.Name, user.Age, user.IsActive)


	celebrateBirthday(&user)

	fmt.Println("\n--- Sau khi tổ chức sinh nhật ---")
	fmt.Printf("Tên: %s, Tuổi: %d, Hoạt động: %t\n", user.Name, user.Age, user.IsActive)


	updateStatus(&user, false)

	fmt.Println("\n--- Sau khi cập nhật trạng thái ---")
	fmt.Printf("Tên: %s, Tuổi: %d, Hoạt động: %t\n", user.Name, user.Age, user.IsActive)
}