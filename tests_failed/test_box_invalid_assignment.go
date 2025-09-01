package main

// This test should fail to compile
// Error: cannot use 3.14 (untyped float constant) as StringOrInt value in assignment

type StringOrInt box {
	string
	int
}

func main() {
	var s StringOrInt
	
	// Valid assignments
	s = "hello"  // OK
	s = 42       // OK
	
	// Invalid assignment - float64 is not a declared variant
	s = 3.14  // COMPILE ERROR: float64 not in box type variants
	
	println(s)
}