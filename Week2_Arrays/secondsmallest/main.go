package main

import (
	"fmt"
	"math"
)

func SmallestAndSecondSmallest(arr []int) []int {
	smallest, secondsmalest := math.MaxInt, math.MaxInt
	for _, v := range arr {
		if v < smallest {
			secondsmalest = smallest
			smallest = v
		} else if v != smallest && v < secondsmalest {
			secondsmalest = v
		}
	}
	if secondsmalest == math.MaxInt {
		return []int{-1}
	}
	return []int{smallest, secondsmalest}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	res := SmallestAndSecondSmallest(arr)
	fmt.Println(res)
}
