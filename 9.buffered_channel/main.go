package main

import "fmt"

func BufferedChannel() {
	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Println("producer done.")
		for i := 0; i < 4; i++ {
			fmt.Printf("sending %d\n", i)
			intStream <- i
		}
	}()

	// for integer := range intStream {
	// 	fmt.Printf("received %d\n", integer)
	// }

	for i := 0; i < 4; i++ {
		fmt.Println(<-intStream)
	}
}

func main() {
	BufferedChannel()
}
