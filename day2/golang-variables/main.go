package main

import (
	"fmt"
	"math"
)

func solveQuadratic(a, b, c float64) {
	// Giải phương trình bậc hai ax^2 + bx + c = 0
	if a == 0 {
		if b == 0 {
			if c == 0 {
				fmt.Println("Phương trình vô số nghiệm")
			} else {
				fmt.Println("Phương trình vô nghiệm")
			}
		} else {
			x := -c / b
			fmt.Printf("Phương trình có một nghiệm: x = %.2f\n", x)
		}
	} else {
		delta := b*b - 4*a*c
		if delta < 0 {
			fmt.Println("Phương trình vô nghiệm")
		} else if delta == 0 {
			x := -b / (2 * a)
			fmt.Printf("Phương trình có nghiệm kép: x1 = x2 = %.2f\n", x)
		} else {
			sqrtDelta := math.Sqrt(delta)
			x1 := (-b + sqrtDelta) / (2 * a)
			x2 := (-b - sqrtDelta) / (2 * a)
			fmt.Printf("Phương trình có hai nghiệm phân biệt: x1 = %.2f, x2 = %.2f\n", x1, x2)
		}
	}
}

func main() {
	var a, b, c float64

	fmt.Println("-- Giai phuong trinh bac hai ax^2 + bx + c = 0 --")

	fmt.Print("Nhap a: ")
	fmt.Scanln(&a)

	fmt.Print("Nhap b: ")
	fmt.Scanln(&b)

	fmt.Print("Nhap c: ")
	fmt.Scanln(&c)

	solveQuadratic(a, b, c)
}