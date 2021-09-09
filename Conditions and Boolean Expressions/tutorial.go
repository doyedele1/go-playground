package main

import "fmt"

func main() {
	x1 := 42
	y1 := 43.5

	x2 := 'd'
	y2 := 'D'

	val1 := float64(x1)+1.5 == y1
	val2 := x2 < y2

	fmt.Printf("%t \n", val1)
	fmt.Printf("%t \n", val2)
}
