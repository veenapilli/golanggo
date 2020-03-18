package main

import "fmt"

/**
Map is a refernce type like a slice
Keys and values are of the same type correspondingly
Keys are indexed. So can be iterated
Keys not known at compile time
*/
func main() {
	// colors:= make(map[string]string)
	// map[key]value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		// every pair is followed by ','. Even the last one

	}
	// iterate over map
	printColors(colors)

	// modify map
	colors["red"] = "#aa0000"

	// delete item from map
	delete(colors, "green")

	fmt.Println("\nModified map:")
	fmt.Println(colors)
}

func printColors(m map[string]string) {
	fmt.Println("\nColors in the map:")
	for color, value := range m {
		fmt.Println(color + " " + value)
	}
}
