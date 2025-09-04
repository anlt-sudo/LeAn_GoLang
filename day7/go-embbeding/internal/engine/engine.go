package models

import "fmt"

type Engine struct {
	Power     int
	IsRunning bool
}

func (e *Engine) Start() {
	e.IsRunning = true
	fmt.Println("Engine started...")
}

func (e *Engine) Stop() {
	e.IsRunning = false
	fmt.Println("Engine stopped...")
}