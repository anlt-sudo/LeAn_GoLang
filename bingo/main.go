package main

import (
	"fmt"
	"math/rand"
	"time"
)

// BingoCard là bảng 5x5
type BingoCard [5][5]string

// Tạo bảng Bingo
func NewBingoCard() BingoCard {
	rand.Seed(time.Now().UnixNano())

	// Tạo danh sách số từ 1..75
	nums := make([]int, 75)
	for i := 0; i < 75; i++ {
		nums[i] = i + 1
	}
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	// Lấy 25 số đầu tiên, đưa vào 5x5
	var card BingoCard
	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			card[i][j] = fmt.Sprintf("%2d", nums[k]) // định dạng 2 chữ số
			k++
		}
	}
	// Ô giữa là FREE
	card[2][2] = " F"
	return card
}

// In bảng Bingo
func (b BingoCard) Print() {
	fmt.Println(" B   I   N   G   O")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf(" %2s ", b[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Đánh dấu số
func (b *BingoCard) Mark(num int) {
	s := fmt.Sprintf("%2d", num)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] == s {
				b[i][j] = " X"
			}
		}
	}
}



func main() {
	card := NewBingoCard()
	fmt.Println("🎲 BẢNG BINGO 🎲")
	card.Print()

	// Giả lập random gọi số
	for round := 1; round <= 100; round++ {
		num := rand.Intn(75) + 1
		fmt.Printf("Gọi số: %d\n", num)
		card.Mark(num)
		card.Print()
	}
}
