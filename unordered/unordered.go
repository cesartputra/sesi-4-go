package main

import (
	"fmt"
	"sync"
)

func printData(data []interface{}, iteration int) {
	fmt.Println(data, iteration)
}

func main() {
	var wg sync.WaitGroup

	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(iteration int) {
			printData(data1, iteration)
			wg.Done()
		}(i + 1)

		wg.Add(1)
		go func(iteration int) {
			printData(data2, iteration)
			wg.Done()
		}(i + 1)
	}

	wg.Wait()
}
