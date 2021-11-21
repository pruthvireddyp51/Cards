package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"

	for key, value := range colors {
		fmt.Println(key, value)
	}
	delete(colors, "red")
	fmt.Println(colors)
}

// All keys must be the same type.