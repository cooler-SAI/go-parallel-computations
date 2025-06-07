package main

import (
	"fmt"
	"sync"
)

func hello(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello", name)
}

func main() {
	fmt.Println("hello world")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("hello world NEW")
	}()

	wg.Add(1)

	go hello("Ander", &wg)
	wg.Wait()
	fmt.Println("Done")

}
