package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DownloadFile(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Downloading file %s...\n", fileName)
	downloadTime := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(downloadTime)

	fmt.Printf("Finished downloading %s after %v\n", fileName, downloadTime)
}