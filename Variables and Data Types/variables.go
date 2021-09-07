package main

import "fmt"

func main() {
	// Go is a statically typed language

	// Declare variables
	var name string = "Demilade"
	var number1 uint8 = 6 // an explicit variable declaration
	number1 = number1 + 17
	fmt.Println(name, number1)

	var number2 = 16             // an implicit variable declaration
	fmt.Printf("%T \n", number2) // %T prints the type of the variable

	number3 := 26 // another way to declare a variable using the expression assignment operator - the walrus operator
	fmt.Printf("%T \n", number3)

	// default values
	var number4 uint64 // is zero for numeric types
	var bl bool        // is false for boolean type
	fmt.Println(number4, bl)
}
