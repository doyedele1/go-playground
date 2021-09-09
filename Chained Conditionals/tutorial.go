package main

import "fmt"

func main() {
	x := 8
	y := 9

	value1 := (true || false) && !false || (x == y)
	value2 := value1 || false
	fmt.Printf("%t", value2)
}
