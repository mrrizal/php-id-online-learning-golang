package main

import (
	"fmt"
	"sync"
)

func hello(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello %s\n", name)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go hello("Rizal", &wg)
	wg.Wait()
}
