package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type your birth year: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) // arguments - the string about to be converted to int, base, bit size
	fmt.Printf("You will be %d years old at the end of 2021", 2021-input)
}
