package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(ch chan int) {
		fmt.Println("start")
		fmt.Println(<-ch)
	}(ch)

	go func() {
		for i := 0; i <= 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("_")
		}
	}()
	time.Sleep(2500 * time.Millisecond)
	ch <- 5

	time.Sleep(3 * time.Second)
}
