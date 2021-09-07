package main

import "fmt"

func main() {
	fmt.Printf("Hello %t \n", true)

	fmt.Printf("Number in base 2: %b \n", 255)

	fmt.Printf("Number: %e \n", 2.2345666694959)

	var out string = fmt.Sprintf("Number: \t %07d is cool", 34)
	fmt.Println(out)
}
