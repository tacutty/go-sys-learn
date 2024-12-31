package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		v := "Hello"
		ch1 <- v
		v2 := <-ch2
		fmt.Println("Goroutine 1:", v, v2)
	}()

	v := "World"
	var v2 string
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println("Main:", v, v2)
}
