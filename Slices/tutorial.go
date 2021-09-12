package main

import "fmt"

func main() {
	/* Slices are an addition to arrays and they are different data types in Go. Slices help to fix
	some problems with standard arrays like determining the size of an array
	when they are created.
	Pointer, length and capacity are all part of the slice.
	In arrays, length and capacity are the same. In slices, length is the size
	of the sliced array while capacity is the maximum size
	that can be stored in the slice.
	*/

	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = x[1:3]
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// fmt.Println(s[:cap(s)])
	// fmt.Println(s[1:])

	// var a []int = []int{5, 6, 7, 8, 9} // creating a slice
	// fmt.Println(cap(a[:3]))
	// a = append(a, 10) // return a new slice with 10 appended to the end of the old slice
	// fmt.Println(a)

	// Another way to make a slice. This would help in concurrency
	a := make([]int, 5)
	fmt.Printf("%T \n", a)
}
