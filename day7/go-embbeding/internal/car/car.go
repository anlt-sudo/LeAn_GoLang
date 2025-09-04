package car

import (
	"fmt"
	models "go-embbeding/internal/engine"
)

type Car struct {
	Brand string
	*models.Engine
}

func (c *Car) Start(){
	fmt.Println("Car is starting...")
	c.Engine.Start()
}