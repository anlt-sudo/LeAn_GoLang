package service

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type FetchResult struct {
	URL   string
	Title string
	Error error
}

func FetchURLTitle(url string) (string, error) {
	fmt.Printf("Fetching %s...\n", url);

	time.Sleep(time.Duration(50 + rand.Intn(450)) * time.Microsecond)

	if strings.Contains(url, "error"){
		return "", fmt.Errorf("Failed to fetch %s", url)
	}

	return fmt.Sprintf("Title for %s", url), nil
}

func FetchAndReport(url string, resultsChan chan <- FetchResult, wg *sync.WaitGroup){
	defer wg.Done()

	title, err := FetchURLTitle(url)

	resultsChan <- FetchResult{URL : url, Title: title, Error: err}
}