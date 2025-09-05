package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"go-channel-use-select/internal/service"
)



func main() {
	rand.Seed(time.Now().UnixNano())

	bidsChan := make(chan service.Bid)
	var wg sync.WaitGroup

	go service.Auctioneer(bidsChan, 5*time.Second)

	bidders := []string{"Alice", "Bob", "Charlie"}
	for _, name := range bidders {
		wg.Add(1)
		go service.Bidder(name, bidsChan, &wg)
	}

	wg.Wait()
	time.Sleep(6 * time.Second)
	fmt.Println("Chương trình chính kết thúc.")
}