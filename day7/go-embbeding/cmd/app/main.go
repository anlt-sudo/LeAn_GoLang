package main

import (
	"fmt"
	"go-embbeding/internal/car"
	models "go-embbeding/internal/engine"
)

func main() {
	engineV8 := &models.Engine{
		Power: 450,
		IsRunning: false,
	}

	mustang := car.Car{
		Brand:  "Ford Mustang",
		Engine: engineV8,
	}

	fmt.Println("Car brand: ", mustang.Brand)
	fmt.Println("Let's start the car...")
	mustang.Start()
	fmt.Println("Is the engine running? ", mustang.IsRunning)
	fmt.Println("Stopping the engine...")
	mustang.Stop()
	fmt.Println("Is the engine running? ", mustang.IsRunning)

}