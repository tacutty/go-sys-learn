package main

import (
	"fmt"
	"os"
)

func main() {
	// Set environment variable
	os.Setenv("NAME", "Gopher")

	// Get environment variable
	name := os.Getenv("NAME")
	fmt.Printf("Hello %s!\n", name)

	// Expand environment variable
	msg := os.ExpandEnv("Hello $NAME!")
	fmt.Println(msg)
}