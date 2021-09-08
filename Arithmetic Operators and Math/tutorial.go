package main

import (
	"fmt"
	"math"
)

func main() {
	var number1 float64 = 9
	var number2 float64 = 4

	var number = math.Sqrt(number1)

	answer := number / number2 // the two operands need to be of the same type

	fmt.Printf("%g", answer)
}
