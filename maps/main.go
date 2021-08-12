package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#00000",
		"white": "#fffff",
	}
	// colors := make(map[int]string)
	// colors[10] = "#fffff"
	// delete(colors, 10)

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, hex)
	}
}
