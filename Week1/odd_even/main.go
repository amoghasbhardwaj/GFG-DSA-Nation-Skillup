package main

import "fmt"

func oddeven(n int) bool {
	return n%2 == 0
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(oddeven(n))
}
