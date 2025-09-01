package main

import "fmt"

// API Result pattern - simplified to avoid compiler issue
type APIResult box {
	string  // Error message
	int     // Status code
}

func fetchData(id int) APIResult {
	if id < 0 {
		return "Invalid ID: negative id"
	}
	if id == 0 {
		return 404  // Not found status
	}
	return 200  // Success status
}

// Configuration pattern
type HostList struct {
	Hosts []string
}

type ConfigValue box {
	string
	int
	bool
	HostList
}

func getConfig(key string) ConfigValue {
	switch key {
	case "name":
		return "MyApp"
	case "port":
		return 8080
	case "debug":
		return false
	case "hosts":
		return HostList{Hosts: []string{"localhost", "127.0.0.1"}}
	default:
		return "unknown"
	}
}

// Resource state pattern
type Uninitialized struct{}
type Initializing struct {
	Progress int
}
type Ready struct {
	ID   string
	Name string
}
type Failed struct {
	Reason string
}

type ResourceState box {
	Uninitialized
	Initializing
	Ready
	Failed
}

func main() {
	// API Result pattern
	fmt.Println("API Result Pattern:")
	results := []APIResult{
		fetchData(1),
		fetchData(0),
		fetchData(-1),
	}
	
	for _, result := range results {
		switch r := result.(type) {
		case string:
			fmt.Printf("Error: %s\n", r)
		case int:
			if r == 200 {
				fmt.Printf("Success: status %d\n", r)
			} else {
				fmt.Printf("Error: status %d\n", r)
			}
		}
	}
	
	// Configuration pattern
	fmt.Println("\nConfiguration Pattern:")
	configs := []string{"name", "port", "debug", "hosts", "unknown"}
	
	for _, key := range configs {
		value := getConfig(key)
		fmt.Printf("%s = ", key)
		switch v := value.(type) {
		case string:
			fmt.Printf("%s (string)\n", v)
		case int:
			fmt.Printf("%d (int)\n", v)
		case bool:
			fmt.Printf("%v (bool)\n", v)
		case HostList:
			fmt.Printf("%v (HostList)\n", v.Hosts)
		}
	}
	
	// Resource state pattern
	fmt.Println("\nResource State Pattern:")
	var resource ResourceState = Uninitialized{}
	
	switch resource.(type) {
	case Uninitialized:
		fmt.Println("Resource is uninitialized")
	case Initializing:
		fmt.Println("Resource is initializing")
	case Ready:
		fmt.Println("Resource is ready")
	case Failed:
		fmt.Println("Resource failed")
	}
}