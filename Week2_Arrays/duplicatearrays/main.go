package duplicatearrays

import "fmt"

func FindDuplicates(arr []int) []int {
	var res []int

	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if val < 0 { // if already negative, get its positive value
			val = -val
		}

		index := val - 1 // map value 1..n → index 0..n-1

		if arr[index] < 0 {
			// already visited → duplicate found
			res = append(res, val)
		} else {
			// mark as visited by making negative
			arr[index] = -arr[index]
		}
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	res := FindDuplicates(arr)
	fmt.Println(res)
}
