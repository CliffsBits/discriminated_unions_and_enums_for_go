package main

import "fmt"

// Result type for safe returns
type Result box {
	int
	string
}

// Divide with single return value
func divide(a, b int) Result {
	if b == 0 {
		return "division by zero"
	}
	return a / b
}

// Optional type
type None struct{}

type Optional box {
	int
	None
}

func findValue(values []int, target int) Optional {
	for _, v := range values {
		if v == target {
			return v
		}
	}
	return None{}
}

func main() {
	// Test divide function
	result1 := divide(10, 2)
	switch val := result1.(type) {
	case int:
		fmt.Printf("Result: %d\n", val)
	case string:
		fmt.Printf("Error: %s\n", val)
	}
	
	result2 := divide(10, 0)
	switch val := result2.(type) {
	case int:
		fmt.Printf("Result: %d\n", val)
	case string:
		fmt.Printf("Error: %s\n", val)
	}
	
	// Test optional returns
	values := []int{1, 2, 3, 4, 5}
	
	opt1 := findValue(values, 3)
	switch val := opt1.(type) {
	case int:
		fmt.Printf("Found value: %d\n", val)
	case None:
		fmt.Println("Value not found")
	}
	
	opt2 := findValue(values, 10)
	switch val := opt2.(type) {
	case int:
		fmt.Printf("Found value: %d\n", val)
	case None:
		fmt.Println("Value not found")
	}
}