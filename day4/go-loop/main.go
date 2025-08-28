package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkRevertString(s1, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 || len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}

func checkSymmetricalString(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Kiểm tra Palindrome")
		fmt.Println("2. Kiểm tra chuỗi đảo ngược")
		fmt.Println("0. Thoát")
		fmt.Print("Chọn chức năng: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Print("Nhập chuỗi: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if checkSymmetricalString(text) {
				fmt.Println("✅ Chuỗi này là palindrome.")
			} else {
				fmt.Println("❌ Chuỗi này không phải palindrome.")
			}

		case "2":
			fmt.Print("Nhập chuỗi 1: ")
			text_1, _ := reader.ReadString('\n')
			fmt.Print("Nhập chuỗi 2: ")
			text_2, _ := reader.ReadString('\n')
			text_1 = strings.TrimSpace(text_1)
			text_2 = strings.TrimSpace(text_2)

			if checkRevertString(text_1, text_2) {
				fmt.Println("👉 Oke, đó là chuỗi đảo ngược!")
			} else {
				fmt.Println("👉 Không phải là chuỗi đảo ngược đâu!")
			}


		case "0":
			fmt.Println("👋 Thoát chương trình...")
			return

		default:
			fmt.Println("⚠️ Lựa chọn không hợp lệ, vui lòng chọn lại!")
		}
	}
}
