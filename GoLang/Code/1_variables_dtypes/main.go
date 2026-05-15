package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	var num_int int32 = 22
	var num_flt float32 = 22.5

	fmt.Printf("The value of num_int is %d and the value of num_flt is %f\n", num_int, num_flt)

	// 1. Can't perform operations between different data types without type conversion
	var sum int32 = num_int + int32(num_flt) // Type conversion from float32 to int32
	fmt.Printf("The sum of num_int and num_flt is %d\n", sum)

	// 2. Strings
	var name string = "Alice"
	var name_joined string = name + " Smith" // String concatenation
	var name_length int = len(name)          // Length of the string
	var multiline_string string = `This is a
multiline string`
	fmt.Printf("Name: %s, Joined Name: %s, Multiline : %s, Name Length: %d\n", name, name_joined, multiline_string, name_length)

	// 3. Booleans
	var cond bool = true
	fmt.Printf("The value of cond is %t\n", cond)

	// 4. Type Inference
	var num2 = 10        // Go infers this as int
	var mark = 85.5      // Go infers this as float64
	var name_2 = "Alice" // Go infers this as string
	var cond2 = false    // Go infers this as bool

	fmt.Printf("num2: %d, mark: %f, name: %s, cond2: %t\n", num2, mark, name_2, cond2)

	// 5. Walrus Operator (Go doesn't have a walrus operator, but we can achieve similar functionality with short variable declaration)
	name3 := "hello"     // Short variable declaration
	val1, val2 := 10, 20 // Multiple variable declaration with short variable declaration

	fmt.Printf("name3: %s, val1: %d, val2: %d\n", name3, val1, val2)

	// 6. Constants
	const PI float64 = 3.14159
	fmt.Printf("The value of PI is %f\n", PI)
}
