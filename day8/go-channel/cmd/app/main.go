package main

import (
	"fmt"
	"go-channel/internal/service"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	urls := []string {
		"https://example.com/1",
		"https://example.com/2",
		"https://example.com/error-page",
		"https://example.com/3",
		"https://example.com/4",
		"https://example.com/5",
	}

	var wg sync.WaitGroup

	// Created a buffered channel to hold results
	resultsChan := make(chan service.FetchResult, len(urls))

	fmt.Println("Starting fetches...")

	for _, url := range urls{
		wg.Add(1)
		go service.FetchAndReport(url, resultsChan, &wg)
	}

	go func ()  {
		wg.Wait()
		close(resultsChan)
	}()

	finalResults := make(map[string]string)

	for result := range resultsChan{
		if result.Error != nil{
			fmt.Printf("Error processing %s: %v\n", result.URL, result.Error)
		}else{
			finalResults[result.URL] = result.Title
		}
	}

	fmt.Println("\n--- Final Results ---")
	for url, title := range finalResults{
		fmt.Printf("%s -> %s\n", url, title)
	}
}