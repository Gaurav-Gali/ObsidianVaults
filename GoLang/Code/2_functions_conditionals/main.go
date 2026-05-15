package main

import (
	"errors"
	"fmt"
)

func main() {
	num1 := 10
	num2 := 2
	q, r, err := division(num1, num2)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else if r == 0 {
		fmt.Printf("Numerator : %d, Denominator : %d\nQuotient : %d\n", num1, num2, q)
	} else {
		fmt.Printf("Numerator : %d, Denominator : %d\nQuotient : %d, Remainder : %d\n", num1, num2, q, r)
	}

	// with switch case
	switch err {
	case nil:
		fmt.Printf("Numerator : %d, Denominator : %d\nQuotient : %d, Remainder : %d\n", num1, num2, q, r)
	default:
		fmt.Printf("Error: %s\n", err.Error())

	}
}

func division(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("Denominator cannot be zero")
		return 0, 0, err
	}

	q := numerator / denominator
	r := numerator % denominator
	return q, r, err
}
