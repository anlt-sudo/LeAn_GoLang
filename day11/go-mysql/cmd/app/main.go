package main

import (
	"bufio"
	"fmt"
	"go-mysql/internal/config"
	"go-mysql/internal/model"
	"go-mysql/internal/repository"
	"os"
	"strconv"
	"strings"
)

func main() {
	db := config.ConnectDB()

	defer db.Close()

	albumRepo := repository.AlbumRepository{DB: db}
	categoryRepo := repository.CategoryRepository{DB: db}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n==== Album Menu ====")
		fmt.Println("1. Thêm album mới")
		fmt.Println("2. Tìm kiếm album theo tên tác giả")
		fmt.Println("3. Tìm kiếm album theo ID")
		fmt.Println("4. Thêm Danh Mục")
		fmt.Println("5. Lấy Tất cả Danh Mục")
		fmt.Println("6. Lấy Tất cả Album")
		fmt.Println("0. Thoát")
		fmt.Print("Chọn thao tác: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			categories, err := categoryRepo.GetAllCategories()
			if err != nil {
				fmt.Println("Lỗi lấy danh sách danh mục:", err)
				break
			}

			if len(categories) == 0 {
				fmt.Println("Chưa có danh mục nào, vui lòng thêm danh mục trước!")
				break
			}

			fmt.Println("Danh sách danh mục:")
			for _, cat := range categories {
				fmt.Printf("  %s - %s\n", cat.ID, cat.Name)
			}

			fmt.Print("Nhập tên album: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Nhập tên tác giả: ")
			artist, _ := reader.ReadString('\n')
			artist = strings.TrimSpace(artist)
			fmt.Print("Nhập giá: ")
			priceStr, _ := reader.ReadString('\n')
			priceStr = strings.TrimSpace(priceStr)
			price, _ := strconv.ParseFloat(priceStr, 64)
			fmt.Print("Nhập mã danh mục (chọn từ trên): ")
			catID, _ := reader.ReadString('\n')
			catID = strings.TrimSpace(catID)
			newID, err := albumRepo.AddAlbum(model.Album{
				Title:  title,
				Artist: artist,
				Price:  float32(price),
				CategoryID: catID,
			})
			if err != nil {
				fmt.Println("Lỗi thêm album:", err)
			} else {
				fmt.Printf("Thêm album thành công! ID: %s\n", newID)
			}
		case "2":
			fmt.Print("Nhập tên tác giả: ")
			artist, _ := reader.ReadString('\n')
			artist = strings.TrimSpace(artist)
			albums, err := albumRepo.AlbumsByArtist(artist)
			if err != nil {
				fmt.Println("Lỗi tìm kiếm:", err)
			} else {
				fmt.Printf("Kết quả tìm kiếm (%d album):\n", len(albums))
				for _, a := range albums {
					fmt.Printf("- ID: %s | %s - %s (%.2f)\n", a.ID, a.Title, a.Artist, a.Price)
				}
			}
		case "3":
			fmt.Print("Nhập ID album: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			album, err := albumRepo.AlbumByID(idStr)
			if err != nil {
				fmt.Println("Lỗi tìm kiếm:", err)
			} else {
				fmt.Printf("Album: ID: %s | %s - %s (%.2f$) - %s\n", album.ID, album.Title, album.Artist, album.Price, album.CategoryID)
			}
		case "4":
			fmt.Print("Nhập tên danh mục mới: ")
			catName, _ := reader.ReadString('\n')
			catName = strings.TrimSpace(catName)

			newCatID, err := categoryRepo.AddCategory(catName)
			if err != nil {
				fmt.Println("Lỗi thêm danh mục:", err)
			} else {
				fmt.Printf("Thêm danh mục thành công! ID: %s\n", newCatID)
			}

		case "5":
			categories, err := categoryRepo.GetAllCategories()
			if err != nil {
				fmt.Println("Lỗi lấy danh sách danh mục:", err)
			} else {
				fmt.Println("Danh sách danh mục:")
				for _, c := range categories {
					fmt.Printf("  %s - %s\n", c.ID, c.Name)
				}
			}

		case "6":
			albums, err := albumRepo.GetAllAlbums()
			if err != nil {
				fmt.Println("Lỗi lấy danh sách album:", err)
			} else {
				fmt.Println("Danh sách album:")
				for _, a := range albums {
					fmt.Printf("  %s - %s (%s) | %.2f | Danh mục: %s\n",
						a.ID, a.Title, a.Artist, a.Price, a.CategoryID)
				}
			}

		case "0":
			fmt.Println("Đã thoát!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ!")
		}
	}

}