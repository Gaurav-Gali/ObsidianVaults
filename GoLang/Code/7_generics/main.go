package main

import (
	"fmt"
)

func sum[T int | float32 | float64](arr []T) T {
	var sum T = 0

	for _, v := range arr {
		sum += v
	}

	return sum
}

func main() {
	var arr_int []int = []int{1, 2, 3, 4, 5}
	var arr_flt []float32 = []float32{1.1, 2.2, 3.3, 4.4, 5.5}

	sum_int := sum(arr_int)
	sum_flt := sum(arr_flt)

	fmt.Println("Sum of integer array is ", sum_int)
	fmt.Println("Sum of float array is ", sum_flt)
}
