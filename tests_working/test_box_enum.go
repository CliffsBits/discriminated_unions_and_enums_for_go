package main

import "fmt"

// Define enum variants as named empty structs
type Red struct{}
type Green struct{}
type Blue struct{}

// Box type as enum
type Color box {
	Red
	Green
	Blue
}

func printColor(c Color) {
	switch c.(type) {
	case Red:
		fmt.Println("Color is Red")
	case Green:
		fmt.Println("Color is Green")
	case Blue:
		fmt.Println("Color is Blue")
	// No default - exhaustive checking
	}
}

func main() {
	var color1 Color = Red{}
	var color2 Color = Green{}
	var color3 Color = Blue{}
	
	printColor(color1)
	printColor(color2)
	printColor(color3)
	
	// Pattern matching with type switch
	colors := []Color{Red{}, Green{}, Blue{}, Red{}}
	
	for i, c := range colors {
		fmt.Printf("Color %d: ", i)
		switch c.(type) {
		case Red:
			fmt.Println("Red")
		case Green:
			fmt.Println("Green")
		case Blue:
			fmt.Println("Blue")
		}
	}
}