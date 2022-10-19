package main

import (
	"fmt"
	"sync"
	"time"
)

func mutexExample() {
	type IntData struct {
		value int
		m     sync.Mutex
	}

	var intData IntData
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			intData.m.Lock()
			defer intData.m.Unlock()
			time.Sleep(time.Second)
			fmt.Printf("value: %d\n", intData.value)
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			intData.m.Lock()
			defer intData.m.Unlock()
			time.Sleep(time.Second)
			fmt.Println("incrementing data!")
			intData.value += 1
		}()
	}

	wg.Wait()
}

func rwMutexExample() {
	type IntData struct {
		value int
		m     sync.RWMutex
	}

	var intData IntData
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			intData.m.RLock()
			defer intData.m.RUnlock()
			time.Sleep(time.Second)
			fmt.Printf("value: %d\n", intData.value)
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			intData.m.Lock()
			defer intData.m.Unlock()
			time.Sleep(time.Second)
			fmt.Println("incrementing data!")
			intData.value += 1
		}()
	}

	wg.Wait()
}

func main() {
	mutexExample()
	// rwMutexExample()
}
