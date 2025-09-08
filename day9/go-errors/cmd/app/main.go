package main

import "go-errors/internal/service"

func main() {
	service.Login("user", "1234567")
	service.Login("", "")
	service.Login("longusername", "mypassword")
}
