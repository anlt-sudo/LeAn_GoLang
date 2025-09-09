package internal

import (
	"fmt"
	"os"
)

type FileHandler struct {
	file *os.File
}

// NewFileHandler tạo file handler mới
func NewFileHandler(filename string) (*FileHandler, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo file: %w", err)
	}

	return &FileHandler{
		file: file,
	}, nil
}

// Close đóng file
func (fh *FileHandler) Close() error {
	if fh.file != nil {
		return fh.file.Close()
	}
	return nil
}

// WriteBoard ghi bảng bingo ban đầu vào file
func (fh *FileHandler) WriteBoard(board [][]string) error {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			_, err := fmt.Fprintf(fh.file, "%2s ", board[i][j])
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprintln(fh.file)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteNewline ghi xuống dòng mới
func (fh *FileHandler) WriteNewline() error {
	_, err := fmt.Fprintln(fh.file)
	return err
}

// WriteCalledNumber ghi số được gọi
func (fh *FileHandler) WriteCalledNumber(number string) error {
	_, err := fmt.Fprintf(fh.file, "%s ", number)
	return err
}

// WriteBingoResult ghi kết quả bingo
func (fh *FileHandler) WriteBingoResult(message string) error {
	_, err := fmt.Fprintf(fh.file, "\n%s\n", message)
	return err
}

// WriteFinalBoard ghi bảng cuối cùng với highlight các vị trí bingo
func (fh *FileHandler) WriteFinalBoard(board [][]string, called map[string]bool, bingoPos [][2]int) error {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			val := board[i][j]
			if val != "X" && called[val] {
				val = "0"
			}
			for _, p := range bingoPos {
				if p[0] == i && p[1] == j {
					val = "A"
				}
			}
			_, err := fmt.Fprintf(fh.file, "%2s ", val)
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprintln(fh.file)
		if err != nil {
			return err
		}
	}
	return nil
}
