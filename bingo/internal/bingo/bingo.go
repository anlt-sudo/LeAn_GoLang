package bingo

import (
	"fmt"
	"math/rand"
)

const SIZE = 5

func CreateBingoBoard() [][]string {
	matrix := make([][]string, SIZE)
	used := make(map[int]bool)
	for i := 0; i < SIZE; i++ {
		matrix[i] = make([]string, SIZE)
		for j := 0; j < SIZE; j++ {
			if i == SIZE/2 && j == SIZE/2 {
				matrix[i][j] = "X"
				continue
			}
			for {
				num := rand.Intn(50) + 1
				if !used[num] {
					matrix[i][j] = fmt.Sprintf("%2d", num)
					used[num] = true
					break
				}
			}
		}
	}
	return matrix
}

func CheckBingo(matrix [][]string, used map[string]bool) (bool, string, [][2]int) {
	for i := 0; i < SIZE; i++ {
		ok := true
		pos := [][2]int{}
		for j := 0; j < SIZE; j++ {
			val := matrix[i][j]
			if val != "X" && !used[val] {
				ok = false
				break
			}
			pos = append(pos, [2]int{i, j})
		}
		if ok {
			return true, fmt.Sprintf("Bingo ngang dong %d", i+1), pos
		}
	}

	for j := 0; j < SIZE; j++ {
		ok := true
		pos := [][2]int{}
		for i := 0; i < SIZE; i++ {
			val := matrix[i][j]
			if val != "X" && !used[val] {
				ok = false
				break
			}
			pos = append(pos, [2]int{i, j})
		}
		if ok {
			return true, fmt.Sprintf("Bingo ngang cot %d", j+1), pos
		}
	}

	ok := true
	pos := [][2]int{}
	for i := 0; i < SIZE; i++ {
		val := matrix[i][i]
		if val != "X" && !used[val] {
			ok = false
			break
		}
		pos = append(pos, [2]int{i, i})
	}
	if ok {
		return true, "Bingo duong cheo chinh", pos
	}

	ok = true
	pos = [][2]int{}
	for i := 0; i < SIZE; i++ {
		val := matrix[i][SIZE-i-1]
		if val != "X" && !used[val] {
			ok = false
			break
		}
		pos = append(pos, [2]int{i, SIZE - i - 1})
	}
	if ok {
		return true, "Bingo duong cheo phu", pos
	}

	return false, "", nil
}
