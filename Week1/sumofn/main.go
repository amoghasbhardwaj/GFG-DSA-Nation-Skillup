package main

import "fmt"

func sumofn(n int) int {
	return n * (n + 1) / 2
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(sumofn(n))
}
