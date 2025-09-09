package main

import (
	"fmt"
	"sync"
	"time"
)

func showGreat(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		fmt.Println(str)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go showGreat("GO", &wg)
	go showGreat("LANG", &wg)

	showGreat("This Is Main", &wg)
}
