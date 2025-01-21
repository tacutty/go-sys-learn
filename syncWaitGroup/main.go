package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("仕事１")
		wg.Done()
	}()

	go func() {
		fmt.Println("仕事２")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("終了")
}
