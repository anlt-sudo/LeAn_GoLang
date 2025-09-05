package car

import (
	"fmt"
	models "go-embbeding/internal/engine"
)

type Car struct {
	Brand string
	*models.Engine
}

// This method is using for ovverriding the method of embedded struct
func (c *Car) Start(){
	fmt.Println("Car is starting...")
	c.Engine.Start()
}