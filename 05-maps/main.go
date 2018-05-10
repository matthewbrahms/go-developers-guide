package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	// Add something to map
	// colors["white"] = "#ffffff"

	// Delete key and value from map
	// delete(colors, "white")

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#33cc33",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
