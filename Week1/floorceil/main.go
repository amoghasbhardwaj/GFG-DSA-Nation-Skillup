package main

import (
	"fmt"
	"math"
)

func divfloorCeil(a, b float64) []int {
	q := float64(a) / float64(b)
	floorval := int(math.Floor(q))
	ceilval := int(math.Ceil(q))
	return []int{floorval, ceilval}
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(divfloorCeil(a, b))
}
