package main

import (
	"fmt"
	"go-worker-pool/internal/service"
	"sync"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://example.come",
		"https://httpbin.org/get",
		"https://jsonplaceholder.typicode.com/posts",
		"https://www.geeksforgeeks.org",
	}

	numWorkers := 3
	numJobs := len(urls)
	jobs := make(chan service.Job, numJobs)
	results := make(chan service.Result, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			service.Worker(id, jobs, results)
		}(w)
	}

	for i, url := range urls {
		jobs <- service.Job{ID: i + 1, Data: url}
	}
	close(jobs)


	go func() {
		wg.Wait()
		close(results)
	}()

	service.ShowResult(results)

	fmt.Println("All jobs processed.")
}
