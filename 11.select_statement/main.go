package main

import (
	"fmt"
	"time"
)

func main() {
	var intStream = make(chan int)
	var stringStream = make(chan string)
	var boolStream = make(chan bool)

	nData := 5

	go func() {
		defer close(intStream)
		defer close(stringStream)
		for i := 1; i <= nData; i++ {
			intStream <- i
			stringStream <- fmt.Sprintf("data ke %d", i)
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		defer close(boolStream)
		boolStream <- true
	}()

Loop:
	for {
		select {
		case intData := <-intStream:
			fmt.Println(intData)
		case stringData := <-stringStream:
			fmt.Println(stringData)
		case <-boolStream:
			break Loop
		}

		time.Sleep(500 * time.Millisecond)
	}

}
