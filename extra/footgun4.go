package main

import "fmt"

func main() {
	add := make(chan int)
	for x := 0; x < 10; x++ {
		go func(x int) {
			add <- x
		}(x)
	}
	sum := 0
	for x := range add {
		sum += x
	}
	fmt.Println(sum)
}
