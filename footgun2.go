package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sum := 0
	for x := 0; x < 10; x++ {
		go func(x int) {
			wg.Add(1)
			sum += x
			wg.Done()
		}(x)
	}
	wg.Wait()
	fmt.Println(sum)
}
