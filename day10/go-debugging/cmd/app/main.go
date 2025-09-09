package main

func sum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}



func main() {
	var n int = 15
	tong := sum(n)
	println("Tổng từ 1 đến", n, "là:", tong)

}