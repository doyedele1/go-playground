package main

import "fmt"

func main() {
	name := "demi"
	fmt.Println("Before if")
	if name != "demi" {
		fmt.Println("Inside if")
	}
	fmt.Println("After if")
}
