package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// todo: benerin ini
	for _, word := range []string{"good", "day", "hello"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(word)
		}()
	}
	wg.Wait()
}
