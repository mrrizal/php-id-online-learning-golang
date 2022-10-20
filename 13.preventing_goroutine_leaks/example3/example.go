package example3

import (
	"fmt"
	"time"
)

func Example1() {
	doWork := func(strings <-chan string) {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		<-completed
	}

	doWork(nil)
	fmt.Println("Done")
}

func Example2() {
	doWork := func(done chan interface{}, strings <-chan string) {
		terminated := make(chan interface{})
		go func() {
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
		<-terminated
	}

	done := make(chan interface{})
	doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("canceling goroutine")
		close(done)
	}()

	fmt.Println("Done")
}
