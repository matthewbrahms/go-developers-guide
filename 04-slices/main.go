package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

// This example shows that pointers have a section
// in memory that stores a pointer to head, capacity,
// and length.  When we pass that to functions, it does
// *not* pass the underlying array.
