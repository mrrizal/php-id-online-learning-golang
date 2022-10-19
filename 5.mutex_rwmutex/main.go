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
			// todo: comment lock and unlock,
			// see what happen
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
	/*
		mutex nge lock variable/struct/data
		biar ga bisa di akses sama goroutine yg lain

		biasanya untuk mastiin value dari
		variable tersebut sync ketika ada perubahan data
	*/

	mutexExample()
	// rwMutexExample()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("=============================")
	// 	// mutexExample()
	// 	// rwMutexExample()
	// }
}
