package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Bid struct {
	BidderName string
	Amount     int
}

func Auctioneer(bidsChan <-chan Bid, auctionDuration time.Duration) {
	fmt.Println("--- Phiên đấu giá bắt đầu! ---")
	fmt.Printf("Phiên đấu giá sẽ kết thúc sau %v.\n", auctionDuration)

	var highestBid int
	var highestBidder string

	auctionTimer := time.After(auctionDuration)

	for {
		select {
		case bid, ok := <-bidsChan:
			if !ok {
				fmt.Println("Kênh đặt giá đã đóng. Auctioneer kết thúc.")
				return
			}

			fmt.Printf("[%s] đặt giá %d\n", bid.BidderName, bid.Amount)
			if bid.Amount > highestBid {
				highestBid = bid.Amount
				highestBidder = bid.BidderName
				fmt.Printf(">>> Giá cao nhất mới: %d từ %s!\n", highestBid, highestBidder)
			}

		case <-auctionTimer:
			fmt.Println("\n--- HẾT GIỜ! ---")
			if highestBidder != "" {
				fmt.Printf("Người thắng cuộc là %s với giá %d!\n", highestBidder, highestBid)
			} else {
				fmt.Println("Không có ai đặt giá. Phiên đấu giá kết thúc.")
			}
			return
		}
	}
}

func Bidder(name string, bidsChan chan<- Bid, wg *sync.WaitGroup) {
	defer wg.Done()

	lastBid := 0

	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)

		increment := 10 + rand.Intn(100)
		newBid := lastBid + increment
		lastBid = newBid

		bidsChan <- Bid{BidderName: name, Amount: newBid}
	}
}