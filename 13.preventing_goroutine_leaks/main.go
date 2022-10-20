package main

import (
	"github.com/mrrizal/learn-go-concurrency-pattern/3.preventing_goroutine_leaks/example5"
)

func main() {
	// fmt.Println(runtime.NumGoroutine())
	// example1.Example()
	// fmt.Println("end call example function that contain goroutine")
	// fmt.Println(runtime.NumGoroutine()) // return 2

	// better way to close goroutine (the child goroutine to prevent goroutine leaks)
	// fmt.Println(runtime.NumGoroutine())
	// example2.Example()
	// fmt.Println("end call example function that contain goroutine")
	// fmt.Println(runtime.NumGoroutine()) // return 1

	// deadlock example
	// fmt.Println(runtime.NumGoroutine())
	// example3.Example1()
	// fmt.Println(runtime.NumGoroutine())

	// also deadlock
	// fmt.Println(runtime.NumGoroutine())
	// example3.Example2()
	// fmt.Println(runtime.NumGoroutine())

	// all above example is showing how to handle when the case is the function receive data from channel
	// then process it.

	// the next example will show you how to avoid goroutine leaks when the function is read data from channel.
	// for example, when data not give any data, if you not handle with proper way, your code will be stuck.

	// example4.Example()
	example5.Example()
}
