package main

// This test should fail to compile
// Error: default case not allowed in box type switch

type Status box {
	string
	int
	bool
}

func processStatus(s Status) {
	// Invalid: default case in box type switch
	switch val := s.(type) {
	case string:
		println("String:", val)
	case int:
		println("Int:", val)
	default:  // COMPILE ERROR: default case not allowed - must handle all variants exactly
		println("Other type")
	}
}

func main() {
	var s Status = true
	processStatus(s)
}