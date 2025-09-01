package main

import "fmt"

// This test should fail to compile
// Error: interface types cannot be used as box variants

type Stringer interface {
	String() string
}

// Invalid: trying to use interface as box variant
type Value box {
	string
	int
	Stringer  // COMPILE ERROR: interface types not allowed as box variants
}

func main() {
	var v Value = "hello"
	fmt.Println(v)
}