package main

// This test should fail to compile
// Error: non-exhaustive type switch - missing case for bool

type Value box {
	string
	int
	bool
}

func processValue(v Value) {
	// Non-exhaustive switch - missing bool case
	switch val := v.(type) {
	case string:
		println("String:", val)
	case int:
		println("Int:", val)
	// COMPILE ERROR: missing case for bool variant
	// Box types require exhaustive handling - no default allowed
	}
}

func main() {
	var v Value = true
	processValue(v)
}