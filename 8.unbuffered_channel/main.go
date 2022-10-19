package main

import (
	"fmt"
)

func UnBufferedChannel() {
	// unbuffered channel, begitu ada data di channel, harus segera di konsumsi
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		defer fmt.Println("producer done.")
		for i := 0; i < 5; i++ {
			fmt.Printf("sending %d\n", i)
			intStream <- i
		}
	}()

	// for integer := range intStream {
	// 	fmt.Printf("received %d\n", integer)
	// }

	for i := 0; i < 5; i++ {
		fmt.Printf("received %d\n", <-intStream)
	}
}

func main() {
	UnBufferedChannel()
}
