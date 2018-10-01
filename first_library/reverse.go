// Package first_library contains utility functions for working with GO.
package main

import (

  "fmt"
	
)

// Reverse returns its argument string reversed rune-wise left to right.
func main() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
