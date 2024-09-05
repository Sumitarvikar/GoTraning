package main

import (
	"fmt"
)

type Hello struct {
	Value string
}

func AcceptAnything(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("This is a value of type Integer, %d\n", v)
	case string:
		fmt.Printf("This is a value of type String, %s\n", v)
	case bool:
		fmt.Printf("This is a value of type Boolean, %v\n", v)
	case Hello:
		fmt.Printf("This is a value of type Hello, %s\n", v.Value)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var choice int
	fmt.Println("Enter a number between 1 and 4:")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		AcceptAnything(42)
	case 2:
		AcceptAnything("Hello, Go!")
	case 3:
		AcceptAnything(true)
	case 4:
		AcceptAnything(Hello{Value: "Hello, custom type!"})
	default:
		fmt.Println("Invalid choice")
	}
}
