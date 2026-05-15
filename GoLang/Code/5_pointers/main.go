package main

import (
	"fmt"
)

func sum(arr *[]int) int {
	sum := 0

	for _, v := range *arr {
		sum += v
	}

	return sum
}

func main() {
	var num int32 = 10

	var ptr *int32 = &num

	fmt.Println("Value of num is ", num)
	fmt.Println("Address of num is ", &num)
	fmt.Println("Value of ptr is ", ptr)
	fmt.Println("Value at the address stored in ptr is ", *ptr)

	// Functions and pointers
	arr := []int{1, 2, 3, 4, 5}
	s := sum(&arr)

	fmt.Println("Sum of array is ", s)
}
