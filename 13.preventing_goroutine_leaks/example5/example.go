package example5

import (
	"fmt"
	"math/rand"
	"time"
)

func Example() {
	newRandStream := func(done chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case <-done:
					return
				case randStream <- rand.Int():
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	<-randStream
	time.Sleep(5 * time.Second)
}
