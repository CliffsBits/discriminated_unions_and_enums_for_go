package main

import "fmt"

// NOTE: Some box type switches may have implementation issues
// This demonstrates the syntax and expected behavior

type Value box {
	string
	int
	bool
}

func processValue(v Value) {
	// Exhaustive type switch - no default allowed
	switch val := v.(type) {
	case string:
		fmt.Printf("Processing string: %s\n", val)
	case int:
		fmt.Printf("Processing int: %d\n", val)
	case bool:
		fmt.Printf("Processing bool: %v\n", val)
	// No default case - compiler enforces exhaustive handling
	}
}

func main() {
	var v1 Value = "hello"
	var v2 Value = 42
	var v3 Value = false
	
	processValue(v1)
	processValue(v2)
	processValue(v3)
	
	// Inline type switch
	var data Value = 100
	switch d := data.(type) {
	case string:
		fmt.Printf("String data: %s\n", d)
	case int:
		fmt.Printf("Int data: %d\n", d)
	case bool:
		fmt.Printf("Bool data: %v\n", d)
	}
}