package main

import (
	"fmt"
)

func closestNumber(n, m int) int {
	// Closest multiple less than or equal to n
	x := (n / m) * m

	// Next multiple
	y := x + m

	// Compare distances
	if abs(n-x) < abs(n-y) {
		return x
	} else if abs(n-x) > abs(n-y) {
		return y
	} else {
		// If equal distance, pick one with greater absolute value
		if abs(x) > abs(y) {
			return x
		} else {
			return y
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(closestNumber(n, m))
}
