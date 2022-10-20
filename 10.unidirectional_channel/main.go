package main

import "fmt"

func main() {
	// you can create directional channel with this way
	var dataStream = make(chan interface{})

	// <- {channel}, artinya kita akan menerima value/func dari channel

	var receivedStream <-chan interface{}

	// {channel} <-, artinya kita akan mengirim value/func ke channel tersebut
	var sendStream chan<- interface{}

	receivedStream = dataStream
	sendStream = dataStream

	go func() {
		defer close(sendStream)
		for i := 0; i < 10; i++ {
			sendStream <- i
		}
	}()

	for data := range receivedStream {
		fmt.Printf("received :%d\n", data)
	}

}
