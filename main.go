package main

import (
	"fmt"
	"os"
)

func main() {
	// Program Name is always the first (implicit argument)
	cmd := os.Args[0]

	fmt.Printf("Program Name: %s\n", cmd)

	argCount := len(os.Args[1:])
	fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)
}
