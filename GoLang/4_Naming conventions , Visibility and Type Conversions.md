![Source](https://youtu.be/SQWy_J2txsI?list=PLS1QulWo1RIaRoN4vQQCYHWDuubEU8Vij)

#### Shadowing
```go
package main

import "fmt"

var a int = 10 // Package level initialisation

func main() {
	var a int = 20
	fmt.Println(a)
}
```

**Here** : 
1. When printing the value of 'a' , the output would be 20 and not 10.
2. This is because 'a' is declared in the local scope.

**Note** :
1. If the variable should only be accessed within the current package then use a lowercase letter to name a variable
2. If the variable should be accessed outside the current package then use a uppercase letter to name a variable
```go
package main

var apple string = "apple" // cannot be used outside the package
var Apple string = "Apple" // can be used outside the package
```

**Note** : 
1. The length of the variable signifies the lifetime of the variable.
2. For example 'i' can be used for loops and the variable is short lived.
3. In order for the variable to persists , use defined names like "userName" etc.

#### Naming Conventions
1. For acronyms use camel casing
```go
var siteUrl = "https://www.google.com"
```

#### Type casting
1. Generally
```go
var x float64 = 20.5
y := int(x) // converts the float to int
```

2. To convert to or from strings using ==strconv==
```go
package main

import (
	"fmt"
	"strconv"
)

var (
	a int = 10
	b string = "20"
)

func main() {
	c := strconv.Itoa(a)
	d := strconv.Atoi(b)
}
```

