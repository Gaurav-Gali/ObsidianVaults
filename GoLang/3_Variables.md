![Source](https://youtu.be/Om-4U2a2PYA?list=PLS1QulWo1RIaRoN4vQQCYHWDuubEU8Vij)

 ### Topics
 1. Variable Declaration
 2. Redeclaration and Shadowing
 3. Visibility
 4. Naming Conventions
 5. Type Conversions

Basic Structure of a go program
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
```


### Variable Declaration
```go
// 1. using var keyword
var a int // declaration
a = 10 // initialisation

// 2. using var keyword
var b int = 20 // declaration + initialisation

// 3. without var keyword
c := 30

fmt.Println(a , b , c)
```

**Note** : Use the %T format specifier to print the type of the variables
```go
a := 10
fmt.Printf("%T" , a)
```

**Note** : If the variable is declared within the function , then it is called function level variable declaration.
If the variable is declared outside the function , then it is called package level variable declaration.

==While declaring variables in the package level , the walrus operator cannot be used.==

#### Declaring multiple variables at the same time using var block
```go
var (
	name string = "Gaurav"
	age int = 19
)
```