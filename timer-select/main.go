package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	buf int = runtime.NumCPU()
)

func main() {
	timer := time.After(time.Second * 5)
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 0; i <= 50000; i++ {
			select {
			case <-timer:
				fmt.Println("timer")
				return
			default:
				c <- i
			}
		}
	}()
	for i := range c {
		fmt.Println(i)

	}
}
