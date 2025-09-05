package main

import (
	"go-concerrency-race/internal/model"
	"go-concerrency-race/internal/service"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	movie := &model.Movie{
		MovieName:      "Inception",
		AvailableSeats: 5,
	}
	service := &service.BookingService{
		Movie: movie,
	}

	usernames := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace"}
	for _, username := range usernames {
		wg.Add(1)
		go service.BookSeat(username, &wg)
	}
	wg.Wait()
	println("All booking attempts completed. Final available seats:", movie.AvailableSeats)
}