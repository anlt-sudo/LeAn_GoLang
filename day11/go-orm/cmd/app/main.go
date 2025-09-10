package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"go-orm/internal/config"
	"go-orm/internal/models"
	"go-orm/internal/repositories"
)




func main() {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Category{}, &models.Album{})

	categoryRepo := repositories.CategoryRepository{DB: db}
	albumRepo := repositories.AlbumRepository{DB: db}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== MENU QUẢN LÝ ALBUM =====")
		fmt.Println("1. Thêm Category")
		fmt.Println("2. Xem tất cả Category")
		fmt.Println("3. Thêm Album")
		fmt.Println("4. Xem tất cả Album")
		fmt.Println("5. Xóa Category")
		fmt.Println("6. Xóa Album")
		fmt.Println("0. Thoát")
		fmt.Print("Chọn chức năng: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			fmt.Print("Nhập tên Category: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			id, err := categoryRepo.Create(name)
			if err != nil {
				log.Println("Lỗi thêm category:", err)
			} else {
				fmt.Println("Thêm category thành công! ID:", id)
			}

		case 2:
			categories, err := categoryRepo.GetAll()
			if err != nil {
				log.Println("Lỗi lấy danh sách:", err)
				continue
			}
			fmt.Println("Danh sách Category:")
			for _, c := range categories {
				fmt.Printf("- %s | %s\n", c.ID, c.Name)
			}

		case 3:
			fmt.Print("Nhập tên Album: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			categories, _ := categoryRepo.GetAll()
			if len(categories) == 0 {
				fmt.Println("Chưa có category nào, hãy thêm trước.")
				continue
			}
			fmt.Println("Danh sách Category:")
			for i, c := range categories {
				fmt.Printf("%d. %s (%s)\n", i+1, c.Name, c.ID)
			}
			fmt.Print("Chọn số Category: ")
			catChoiceStr, _ := reader.ReadString('\n')
			catChoice, _ := strconv.Atoi(strings.TrimSpace(catChoiceStr))
			if catChoice < 1 || catChoice > len(categories) {
				fmt.Println("Lựa chọn không hợp lệ")
				continue
			}
			categoryID := categories[catChoice-1].ID

			id, err := albumRepo.Create(name, categoryID)
			if err != nil {
				log.Println("Lỗi thêm album:", err)
			} else {
				fmt.Println("Thêm album thành công! ID:", id)
			}

		case 4:
			albums, err := albumRepo.GetAll()
			if err != nil {
				log.Println("Lỗi lấy danh sách album:", err)
				continue
			}
			fmt.Println("Danh sách Album:")
			for _, a := range albums {
				fmt.Printf("- %s | %s | Category: %s\n", a.ID, a.Name, a.Category.Name)
			}

		case 5:
			fmt.Print("Nhập ID Category cần xóa: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			if err := categoryRepo.Delete(id); err != nil {
				log.Println("Lỗi xóa category:", err)
			} else {
				fmt.Println("Xóa category thành công")
			}

		case 6:
			fmt.Print("Nhập ID Album cần xóa: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			if err := albumRepo.Delete(id); err != nil {
				log.Println("Lỗi xóa album:", err)
			} else {
				fmt.Println("Xóa album thành công")
			}

		case 0:
			fmt.Println("Thoát chương trình.")
			return

		default:
			fmt.Println("Lựa chọn không hợp lệ.")
		}
	}
}