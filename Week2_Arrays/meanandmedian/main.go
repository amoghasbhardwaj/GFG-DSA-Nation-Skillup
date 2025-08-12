package main

import (
	"fmt"
	"sort"
)

func MeanAndMedian(arr []int) (mean, median int) {
	sort.Ints(arr)
	sum := 0

	for _, v := range arr {
		sum += v
	}

	l := len(arr)
	mid := l / 2

	mean = sum / l
	if l%2 == 0 {
		median = (arr[mid-1] + arr[mid]) / 2
	} else {
		median = arr[mid]
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	mean, median := MeanAndMedian(arr)
	fmt.Println(mean, median)
}
