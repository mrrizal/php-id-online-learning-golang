package example2

import (
	"fmt"
	"runtime"
	"time"
)

func Example() {
	doWork := func(done chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			fmt.Printf("calling from doWork func: %d\n", runtime.NumGoroutine())
			defer fmt.Println("doWork Exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("canceling goroutine")
		close(done)
	}()

	// join do work goroutine with main goroutine
	<-terminated
	fmt.Println("Done")
}
