package main

import (
	"fmt"
	"sync"
)

func interface1Ordered(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	data := []interface{}{"coba1", "coba2", "coba3"}
	for i := 1; i <= 3; i++ {
		mutex.Lock()
		fmt.Println(data, i)
		mutex.Unlock()
	}
}

func interface2Ordered(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	data := []interface{}{"bisa1", "bisa2", "bisa3"}
	for i := 1; i <= 3; i++ {
		mutex.Lock()
		fmt.Println(data, i)
		mutex.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(2)

	go interface1Ordered(&wg, &mutex)
	go interface2Ordered(&wg, &mutex)

	wg.Wait()
}
