package main

import "fmt"

func main() {
	//var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	//	colors := make(map[int]string) // creating a map using make()
	// colors[10] = "#fff"
	// delete(colors, 10) // delete a specific key inside a map

	fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
