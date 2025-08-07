package main

import "fmt"

func sumofdigits(n int) int {
	sum := 0
	for n > 0 {
		temp := n % 10
		sum += temp
		n = n / 10

	}
	return sum
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(sumofdigits(n))
}
