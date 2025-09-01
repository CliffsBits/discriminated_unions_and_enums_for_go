# Box Types for Go - Design Experiment

A feasibility study and proof-of-concept implementation of discriminated union types (box types) for the Go programming language.

## What This Is

This is a design experiment and feasibility study that:

- Explores implementation feasibility - Demonstrates that discriminated unions can be implemented in Go's compiler and runtime
- Addresses historical concerns - Responds to concerns raised by the Go core team over the years about union types
- Provides a concrete implementation - Offers advocates and researchers a working prototype to evaluate, test, and reason about
- Documents design decisions - Shows how box types integrate with Go's existing type system without disruption
- Enables experimentation - Allows developers to explore discriminated union patterns in Go

## What This Is NOT

This is not:

- A pull request - This is not intended as a submission to the official Go project
- Production-ready code - This is experimental code for research and evaluation purposes
- An advocacy statement - This respects the Go team's design decisions and timeline
- A criticism of Go - Go is awesome - people keep asking for sum types / discriminated unions / enums - we wanted to know what it would take

## Design Philosophy

We tried to adhere to the following:

- Go's core principles - Simplicity, clarity, and orthogonal design
- Existing architecture - Building on Go's established patterns rather than disrupting them
- Historical considerations - Learning from and addressing past concerns raised by language designers
- Incremental approach - Adding capabilities without changing existing behavior - we believe this implementation DOES NOT BREAK existing code

## Quick Start

### Building the Compiler

```bash
# Requires Go 1.21+ as bootstrap compiler
export GOROOT_BOOTSTRAP="/usr/local/go"  # or your Go installation path

# Build (takes several minutes)
cd go-release-branch.go1.25/src
./make.bash  # Linux/macOS
# or ./make.bat on Windows
```

### Hello Box Types

```go
package main

import "fmt"

type Status box {
    string
    int
    bool
}

func main() {
    var status Status = "ready"
    
    switch val := status.(type) {
    case string:
        fmt.Println("Message:", val)
    case int:
        fmt.Println("Code:", val)  
    case bool:
        fmt.Println("Flag:", val)
    }
}
```

```bash
# Run with the new compiler
./go-release-branch.go1.25/bin/go run hello.go
```

## Key Features Demonstrated

### Basic Syntax

```go
type Name box {
    Type1
    Type2
    Type3
}
```

### Return Safety - Single Value Returns

Box types offer an alternative to Go's multiple return values for error handling:

```go
// Standard Go approach
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Box types approach - single return value
type Result box {
    int
    string
}

func divideBox(a, b int) Result {
    if b == 0 {
        return "division by zero"  // Error case
    }
    return a / b  // Success case
}

// Usage with exhaustive checking
result := divideBox(10, 2)
switch val := result.(type) {
case int:
    fmt.Println("Result:", val)
case string:
    fmt.Println("Error:", val)
}
```

### Enum Patterns with Named Types

```go
// Define enum variants as named empty structs  
type Red struct{}
type Green struct{}
type Blue struct{}

type Color box {
    Red
    Green
    Blue
}

func main() {
    var color Color = Red{}
    
    switch color.(type) {
    case Red:
        println("Red selected")
    case Green:
        println("Green selected")
    case Blue:
        println("Blue selected")
    // No default case allowed - exhaustive checking enforced
    }
}
```

### Core Capabilities

- **Type Safety** - Only declared variants can be assigned
- **Exhaustive Checking** - Compiler enforces handling all cases (no default cases allowed)
- **Extraction Required** - Must extract variant before use (like interface{} type assertions)
- **Closed Variant Set** - Finite, predetermined set of allowed types
- **Return Safety** - Alternative patterns for error handling
- **Memory Efficient** - Same layout as interface{}, no extra overhead

### Important: How Box Types Work

Box types are **discriminated unions**, not interfaces. Key points:

- **Cannot call methods directly** on box types - must extract first via type switch
- **Similar to interface{}** - same extraction pattern: `switch v := box.(type)`
- **Compile-time exhaustive checking** - must handle all variants, no default case
- **Only basic types and named types** allowed as variants (no interfaces with methods)

## Documentation

- [Historical Analysis](objections_solved/OBJECTIONS_ASSESSMENT.md) - How box types address past concerns
- [Working Tests](tests_working/) - Examples demonstrating each feature
- [Type Safety Tests](tests_failed/) - Examples of properly rejected invalid usage

## Research Context

This work builds on years of community discussion about discriminated unions in Go:

- Considers concerns raised by Rob Pike, Russ Cox, Ian Lance Taylor, and Robert Griesemer
- Addresses complexity, type system impact, and philosophical considerations
- Demonstrates solutions through working code rather than theoretical arguments
- Maintains Go's commitment to simplicity and orthogonal design

## Acknowledgments

This work stands on the shoulders of:

- The **Go core team** for creating such an elegant and extensible language design
- **Community members** who have thoughtfully explored these concepts over the years
- **Researchers and advocates** who have kept the discussion alive with respectful persistence

---

*This project is offered in the spirit of constructive contribution to the Go community's ongoing discussions about language evolution.*