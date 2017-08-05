package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for x := 0; x < 10; x++ {
		go func() {
			sum += x
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(sum)
}
