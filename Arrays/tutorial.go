package main

import "fmt"

func main() {
	// An array is a specified-ordered collection of elements (some type) used to store multiple values.
	// Declaring an array with integer numbers and fixed maximum size 5
	// var arr [5]int

	// Initializing an array with specific numbers
	arr := [3]int{4, 5, 6}

	// Accessing array elements: arr[0] = 100
	// arr[4] = 900
	// fmt.Println(len(arr))

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

	arr2D := [2][3]int{{1, 2, 3}, {3, 3, 4}}
	fmt.Println(arr2D[0][2])
}
