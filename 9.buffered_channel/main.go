package main

import (
	"fmt"
	"time"
)

func BufferedChannel() {
	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Println("producer done.")
		for i := 0; i < 10; i++ {
			intStream <- i
			fmt.Printf("sending %d\n", i)
		}
	}()

	// baca data dari channel
	// for integer := range intStream {
	// 	fmt.Printf("received %d\n", integer)
	// }

	// another way baca data dari channel
	// for i := 0; i < 4; i++ {
	// 	fmt.Println(<-intStream)
	// }

	// membuktikan buffered channel hanya bisa menerima data sebanyak
	// kapasitas yg sudah ditentukan
	fmt.Println(<-intStream)
	time.Sleep(time.Second)
}

func main() {
	BufferedChannel()
}
