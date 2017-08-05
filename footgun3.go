package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sum := 0
	wg.Add(10)
	for x := 0; x < 10; x++ {
		go func(x int) {
			sum += x
			wg.Done()
		}(x)
	}
	wg.Wait()
	fmt.Println(sum)
}
