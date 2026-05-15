package main

import (
	"fmt"
)

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func main() {
	rect1 := Rectangle{1, 2}

	ar := rect1.area()

	fmt.Println("Area of rectangle is ", ar)
}
