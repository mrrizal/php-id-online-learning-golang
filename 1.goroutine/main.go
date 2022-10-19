package main

import (
	"fmt"
)

func hello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func main() {
	/*
		async, tidak memblock proses
		selanjutnya
	*/
	hello("Rizal")
	// time.Sleep(1 * time.Second)
}
