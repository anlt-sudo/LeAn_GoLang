package model

import "sync"

type Movie struct {
	MovieName      string
	AvailableSeats int
	Mu             sync.Mutex
}