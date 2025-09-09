package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/anlt-sudo/bingo/internal/service"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	page := service.CreateBingoBoard()
	file, err := os.Create("bingo_output.txt")
	if err != nil {
		fmt.Println("Lỗi tạo file:", err)
		return
	}
	defer file.Close()

	for i := 0; i < service.SIZE; i++ {
		for j := 0; j < service.SIZE; j++ {
			fmt.Fprintf(file, "%2s ", page[i][j])
		}
		fmt.Fprintln(file)
	}

	called := make(map[string]bool)
	calledList := []string{}
	var bingoMsg string
	var bingoPos [][2]int

	for {
		num := rand.Intn(50) + 1
		str := fmt.Sprintf("%d", num)
		if called[str] {
			continue
		}
		called[str] = true
		calledList = append(calledList, str)

		if len(calledList) == 1 {
			fmt.Fprintln(file)
		}
		fmt.Fprintf(file, "%s ", str)

		ok, msg, pos := service.CheckBingo(page, called)
		if ok {
			bingoMsg = msg
			bingoPos = pos
			break
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Fprintf(file, "\n%s\n", bingoMsg)

	for i := 0; i < service.SIZE; i++ {
		for j := 0; j < service.SIZE; j++ {
			val := page[i][j]
			if val != "X" && called[val] {
				val = "0"
			}
			for _, p := range bingoPos {
				if p[0] == i && p[1] == j {
					val = "A"
				}
			}
			fmt.Fprintf(file, "%2s ", val)
		}
		fmt.Fprintln(file)
	}
	fmt.Print("Da Ket Thuc")
}