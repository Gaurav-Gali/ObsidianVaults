package main

import (
	"fmt"
)

func main() {
	// 1. Arrays
	nums := [3]int{1, 2, 3}
	fmt.Printf("Array: %v\n, Indexing[1], %v\n", nums, nums[1])

	// Dynamic initialization of array
	nums2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n, Length: %d\n", nums2, len(nums2))

	// 2. Slice
	slice1 := []int32{1, 2, 3}
	fmt.Printf("Slice: %v\n, Length: %d\n, Capacity: %d\n", slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 4)
	fmt.Printf("Slice: %v\n, Length: %d\n, Capacity: %d\n", slice1, len(slice1), cap(slice1))

	slice2 := []int32{4, 5}
	slice1 = append(slice1, slice2...) // Append another slice to the existing slice (IMP : use ... to unpack the slice)
	fmt.Println(slice1)

	slice3 := make([]int32, 5, 100) // Create a slice of length 5 and capacity 100 (IMP : imporves performance when we know the size of the slice in advance)
	fmt.Println(slice3)

	// 3. Maps
	register := map[string]int{"Alicee": 22, "Bob": 25}
	fmt.Printf("Map: %v\n, Value for key 'Alice': %d\n", register, register["Alice"])

	age_of_alice, ok := register["Alice"] // Check if key exists in the map
	switch ok {
	case true:
		fmt.Printf("Alice's age is %d\n", age_of_alice)
	default:
		fmt.Println("Alice is not registered")
	}

	// 4. Loops

	for name, age := range register {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// 5. For loop
	for i := 0; i < len(nums); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, nums[i])
	}

	// 6. While loop (Go doesn't have a while loop, but we can achieve similar functionality with a for loop)
	i := 0
	for i < 10 {
		fmt.Printf("Value of i: %d\n", i)
		i++
	}
}
