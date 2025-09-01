package main

import "fmt"

// NOTE: Box type switches in functions may have implementation issues
// This demonstrates the syntax for multi-variant box types

// Named struct types for demonstration
type Point struct {
	X, Y int
}

type Person struct {
	Name string
	Age  int
}

// Box type with 6 different variants
type MultiValue box {
	string
	int
	bool
	float64
	Point
	Person
}

func describeValue(v MultiValue) {
	// Exhaustive type switch - must handle all 6 variants
	switch val := v.(type) {
	case string:
		fmt.Printf("String value: %q (length: %d)\n", val, len(val))
	case int:
		fmt.Printf("Integer value: %d (hex: 0x%x)\n", val, val)
	case bool:
		if val {
			fmt.Println("Boolean value: true")
		} else {
			fmt.Println("Boolean value: false")
		}
	case float64:
		fmt.Printf("Float value: %.2f\n", val)
	case Point:
		fmt.Printf("Point: (%d, %d)\n", val.X, val.Y)
	case Person:
		fmt.Printf("Person: %s, age %d\n", val.Name, val.Age)
	}
}

func main() {
	fmt.Println("Multi-Variant Box Type Demo")
	fmt.Println("============================")
	
	// Create instances of each variant
	var v1 MultiValue = "Hello, Box Types!"
	var v2 MultiValue = 42
	var v3 MultiValue = true
	var v4 MultiValue = 3.14159
	var v5 MultiValue = Point{X: 10, Y: 20}
	var v6 MultiValue = Person{Name: "Alice", Age: 30}
	
	// Process each variant
	values := []MultiValue{v1, v2, v3, v4, v5, v6}
	
	for i, value := range values {
		fmt.Printf("\nVariant %d: ", i+1)
		describeValue(value)
	}
	
	// Demonstrate type assertions
	fmt.Println("\nType Assertions:")
	if str, ok := v1.(string); ok {
		fmt.Printf("v1 is a string: %s\n", str)
	}
	
	if num, ok := v2.(int); ok {
		fmt.Printf("v2 is an int: %d\n", num)
	}
	
	if point, ok := v5.(Point); ok {
		fmt.Printf("v5 is a Point at (%d, %d)\n", point.X, point.Y)
	}
	
	// Demonstrate handling unknown value
	fmt.Println("\nHandling dynamic values:")
	dynamicValues := []MultiValue{
		"dynamic string",
		999,
		false,
		2.718,
		Point{X: 5, Y: 15},
		Person{Name: "Bob", Age: 25},
	}
	
	for _, dv := range dynamicValues {
		fmt.Print("Processing: ")
		describeValue(dv)
	}
}