package service

import (
	"go-concerrency-race/internal/model"
	"sync"
	"time"
)

type BookingService struct {
	Movie *model.Movie
}


func (s *BookingService) BookSeat(username string, wg *sync.WaitGroup){
	defer wg.Done()
	s.Movie.Mu.Lock()
	defer s.Movie.Mu.Unlock()
	if s.Movie.AvailableSeats > 0 {
		time.Sleep(10 * time.Millisecond)
		println("User", username, "booked a seat successfully. Remaining seats:", s.Movie.AvailableSeats)
		s.Movie.AvailableSeats -= 1
	} else {
		println("User", username, "failed to book a seat. No seats available.")
	}
}