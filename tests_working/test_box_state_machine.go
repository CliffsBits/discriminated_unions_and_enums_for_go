package main

import "fmt"

// NOTE: Box type switch implementation may have issues with function calls
// This demonstrates the intended syntax and behavior

// State types with different data
type Idle struct{}
type Loading struct {
	Progress int
}
type Success struct {
	Data string
}
type Error struct {
	Message string
}

// State machine using box type
type State box {
	Idle
	Loading
	Success
	Error
}

func processState(s State) {
	switch state := s.(type) {
	case Idle:
		fmt.Println("System is idle")
	case Loading:
		fmt.Printf("Loading... %d%%\n", state.Progress)
	case Success:
		fmt.Printf("Success! Data: %s\n", state.Data)
	case Error:
		fmt.Printf("Error: %s\n", state.Message)
	}
}

func main() {
	// State transitions
	var currentState State = Idle{}
	processState(currentState)
	
	currentState = Loading{Progress: 25}
	processState(currentState)
	
	currentState = Loading{Progress: 75}
	processState(currentState)
	
	currentState = Success{Data: "Operation completed"}
	processState(currentState)
	
	// Error state
	var errorState State = Error{Message: "Connection timeout"}
	processState(errorState)
	
	// State machine simulation
	fmt.Println("\nState machine simulation:")
	states := []State{
		Idle{},
		Loading{Progress: 0},
		Loading{Progress: 50},
		Loading{Progress: 100},
		Success{Data: "File downloaded"},
	}
	
	for _, state := range states {
		processState(state)
	}
}