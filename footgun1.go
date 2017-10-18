package main

import (
	"fmt"
	"time"
)

func main() {
	for x := range []int{0, 1, 2, 3, 4} {
		go func() {
			fmt.Println(x)
		}()
	}
	time.Sleep(1 * time.Second)
}
