package main

import "fmt"

func main() {
	age := 13

	if age >= 18 {
		fmt.Println("You can ride the rollercoaster alone")
	} else if age >= 14 {
		fmt.Println("You can ride the rollercoaster with a parent")
	} else {
		fmt.Println("You can't ride the rollercoaster")
		fmt.Printf("Wait %d years", 18-age)
	}
}
