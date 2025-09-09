package main

import (
	"bufio"
	"fmt"
	"go-logging/internal/logdemo"
	"os"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n==== Go Logging Demo ====")
		fmt.Println("1. Log với zap")
		fmt.Println("2. Log với zerolog")
		fmt.Println("0. Thoát")
		fmt.Print("Chọn logger (1/2/0): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			logdemo.LogWithZap()
		case "2":
			logdemo.LogWithZerolog()
		case "0":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ!")
		}
	}
	}