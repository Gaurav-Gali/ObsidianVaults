package main

import (
	"bookings-app/arith"
	"fmt"
)

func main() {
	// 1.Go is a statically typed language.
	// 2. The types have to be statically mentioned
	// 3. But when you assign values to variables, the type of the variable will be determined by the value assigned.
	var a  = 10
	const b string = "Hello"
	fmt.Println(a)
	fmt.Println(b)

	// 4. Data Types
	// string , int , uint8 , uint16 , uint32 , uint64 , float32 , float64 , bool , complex64 , complex128

	// 5. User Inputs
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Your name is:", name)

	// 6. Arrays
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println(arr)

	// 7. Slices
	// 8. Slices are arrays of dynamic size.
	var s []int
	s = append(s, 10)
	s = append(s, 20)
	s = append(s, 30)
	fmt.Println(s)

	// 9. Loops
	for i:=0;i<10;i++ {
		fmt.Println(i)
	}

	// 10. While loop
	k := 1

	for k < 5 {
		fmt.Println(k)
		k++
	} 

	// 11. Infinite Loop
	for {
		fmt.Println("Infinite Loop")
		break
	}

	// 12. Looping over arrays and slices
	arr = [3]int{10, 20, 30}
	for index, value := range arr {
		fmt.Printf("%d. %d\n",index+1, value)
	}

	// 13. Conditional Statements
	num := 5
	if num > 10 {
		fmt.Println(num, "is greater than 10")
	}else{
		fmt.Println(num, "is not greater than 10")
	}

	// 14. Switch Statement
	day := "Monday"
	switch day {
	case "Sunday":
		fmt.Println("It's Sunday")
	case "Saturday":
		fmt.Println("It's Saturday")
	default:
		fmt.Println("It's a weekday")
	}

	// 15. Functions from other package
	arith.Add(10,20)
	arith.Subtract(10,20)

	// 16. Maps
	var myMap  = make(map[string]string)
	myMap["name"] = "John"
	fmt.Println(myMap["name"])

	// 17. Structs
	p := Person{Name:"Gaurav", Age:19}
	fmt.Println(p.Name, p.Age)
	p.Greet()

	// 18. Anonymous Struct
	person := struct {
		Name string
	}{
		Name: "Gaurav",
	}

	fmt.Println(person.Name)

	// 19. Go routines
	go count(5, "1")
	go count(10, "2")

}

func count(num int, msg string) {
	for i:=0;i<num;i++ {
		fmt.Println(msg," : ",num)
	}
}

// Defining Structs
type Person struct {
	Name string
	Age uint16
}

// Functions of Structs
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// The `strings` package in Go provides a collection of functions for manipulating UTF-8 encoded strings. It is part of Go's standard library and offers various utilities for searching, replacing, splitting, and modifying strings.

// ## Importing the Package
// ```go
// import "strings"
// ```

// ## Commonly Used Functions

// ### 1. **Contains** – Check if a substring exists
// ```go
// fmt.Println(strings.Contains("hello world", "world")) // true
// ```

// ### 2. **HasPrefix & HasSuffix** – Check if a string starts or ends with a specific substring
// ```go
// fmt.Println(strings.HasPrefix("golang", "go"))  // true
// fmt.Println(strings.HasSuffix("golang", "lang")) // true
// ```

// ### 3. **Index & LastIndex** – Find the index of a substring
// ```go
// fmt.Println(strings.Index("hello world", "world"))      // 6
// fmt.Println(strings.LastIndex("hello hello", "hello")) // 6
// ```

// ### 4. **Replace & ReplaceAll** – Replace occurrences of a substring
// ```go
// fmt.Println(strings.Replace("foo bar foo", "foo", "baz", 1)) // "baz bar foo"
// fmt.Println(strings.ReplaceAll("foo bar foo", "foo", "baz")) // "baz bar baz"
// ```

// ### 5. **Split & Join** – Split a string into a slice and join a slice into a string
// ```go
// fmt.Println(strings.Split("a,b,c", ",")) // ["a" "b" "c"]
// fmt.Println(strings.Join([]string{"a", "b", "c"}, "-")) // "a-b-c"
// ```

// ### 6. **ToLower & ToUpper** – Convert case
// ```go
// fmt.Println(strings.ToLower("GoLang")) // "golang"
// fmt.Println(strings.ToUpper("GoLang")) // "GOLANG"
// ```

// ### 7. **Trim, TrimSpace, TrimPrefix, TrimSuffix** – Remove leading/trailing characters
// ```go
// fmt.Println(strings.Trim("  hello  ", " "))      // "hello"
// fmt.Println(strings.TrimSpace("  hello  "))      // "hello"
// fmt.Println(strings.TrimPrefix("golang", "go")) // "lang"
// fmt.Println(strings.TrimSuffix("golang", "lang")) // "go"
// ```

// ### 8. **Repeat** – Repeat a string multiple times
// ```go
// fmt.Println(strings.Repeat("go", 3)) // "gogogo"
// ```

// ### 9. **Fields** – Split a string by whitespace
// ```go
// fmt.Println(strings.Fields("hello  world")) // ["hello" "world"]
// ```

// ## Example Program
// ```go
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str := "Hello, World!"
// 	fmt.Println(strings.ToUpper(str))      // "HELLO, WORLD!"
// 	fmt.Println(strings.Contains(str, "Go")) // false
// 	fmt.Println(strings.Split(str, ", "))  // ["Hello" "World!"]
// }
// ```

// The `strings` package is optimized for performance and should be preferred over manual string operations. Let me know if you need more details! 🚀