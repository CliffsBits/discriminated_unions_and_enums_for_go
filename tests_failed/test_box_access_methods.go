package main

// This test should fail to compile
// Error: cannot convert b (variable of box type StringOrInt) to type string

type StringOrInt box {
	string
	int
}

func main() {
	var b StringOrInt = "hello"
	
	// Invalid direct casting - must use type assertion
	var s string = string(b)  // COMPILE ERROR: cannot convert box to string directly
	
	// Correct way would be:
	// s, ok := b.(string)
	
	println(s)
}