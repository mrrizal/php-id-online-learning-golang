package example4

import (
	"fmt"
	"math/rand"
	"time"
)

func Example() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			// line 14 never called
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	time.Sleep(3 * time.Second)
	// the newRandStream stop not because it's done, it's because the main program is ended
	// in this case, we simulate Example func took 3 second to proccess.
}
