package main

import "fmt"

type Response box {
	string
	int
	bool
}

func main() {
	var r1 Response = "success"
	var r2 Response = 200
	var r3 Response = true
	
	// Type assertions with ok pattern
	if str, ok := r1.(string); ok {
		fmt.Printf("String response: %s\n", str)
	}
	
	if code, ok := r2.(int); ok {
		fmt.Printf("Int response: %d\n", code)
	}
	
	if flag, ok := r3.(bool); ok {
		fmt.Printf("Bool response: %v\n", flag)
	}
	
	// Check for wrong type assertion
	if _, ok := r1.(int); !ok {
		fmt.Println("r1 is not an int (expected)")
	}
	
	// Direct assertion (panics if wrong type)
	strVal := r1.(string)
	fmt.Printf("Direct assertion: %s\n", strVal)
}