package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := map[string]string{}
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
	}

	colors["blue"] = "#0000ff"
	delete(colors, "red")

	printMap(colors)

	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Hex code for", k, "is", v)
	}
}
