package main

import "fmt"

// NOTE: fmt/string support for boxes is not fully implemented yet
// Box values must be unpacked before printing/string operations

// Basic box type with string and int variants
type Status box {
	string
	int
}

func main() {
	// Assign string variant
	var s1 Status = "ready"
	// Must unpack to print
	if str, ok := s1.(string); ok {
		fmt.Printf("String status: %s\n", str)
	}
	
	// Assign int variant
	var s2 Status = 404
	// Must unpack to print
	if num, ok := s2.(int); ok {
		fmt.Printf("Int status: %d\n", num)
	}
	
	// Type assertion to access value
	if str, ok := s1.(string); ok {
		fmt.Printf("Extracted string: %s\n", str)
	}
	
	if code, ok := s2.(int); ok {
		fmt.Printf("Extracted int: %d\n", code)
	}
}