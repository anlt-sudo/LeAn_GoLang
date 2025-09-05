package main

import (
	"fmt"
	"go-conccurency/internal/service"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	files := []string{
		"image1.jpg",
		"document.pdf",
		"archive.zip",
		"video.mp4",
		"song.mp3",
	}
	var wg sync.WaitGroup

	fmt.Println("Starting parallel download...")
	startTime := time.Now()

	for _, file := range files {
		wg.Add(1)
		go service.DownloadFile(file, &wg)
	}

	fmt.Println("Main: Waiting for downloads to complete...")
	wg.Wait()

	duration := time.Since(startTime)
	fmt.Printf("\nAll files have been downloaded successfully in %v.\n", duration)
}