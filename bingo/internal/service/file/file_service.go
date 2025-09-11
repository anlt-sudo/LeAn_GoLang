package file

import (
	"fmt"
	"os"
)

type FileService struct {
	file *os.File
}

func NewFileService(filename string) (*FileService, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("lỗi tạo file: %w", err)
	}

	return &FileService{
		file: file,
	}, nil
}

func (fh *FileService) Close() error {
	if fh.file != nil {
		return fh.file.Close()
	}
	return nil
}

func (fh *FileService) WriteBoard(board [][]string) error {
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

func (fh *FileService) WriteNewline() error {
	_, err := fmt.Fprintln(fh.file)
	return err
}

func (fh *FileService) WriteCalledNumber(number string) error {
	_, err := fmt.Fprintf(fh.file, "%s ", number)
	return err
}

func (fh *FileService) WriteBingoResult(message string) error {
	_, err := fmt.Fprintf(fh.file, "\n%s\n", message)
	return err
}

func (fh *FileService) WriteFinalBoard(board [][]string, called map[string]bool, bingoPos [][2]int) error {
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
