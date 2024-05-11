package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf854",
		"white": "#ffffff",
	}

	// colors["white"] = "#fffff"

	// delete(colors, "white")

	printMap(colors)

}
