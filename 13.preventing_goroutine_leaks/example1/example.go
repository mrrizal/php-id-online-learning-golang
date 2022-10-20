package example1

import "fmt"

func Example() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	terminated := doWork(nil)
	// join do work goroutine with main goroutine
	// will cause deadlock
	<-terminated
	fmt.Println("Done")
}
