package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	buf := make(chan int)

	n := func() int {
		i := 0
		for ; i <= 100; i++ {
			i++
		}
		return i
	}()

	fmt.Println("100:", n)

	num := []int{1, 2, 3, 4, 5}
	go func() {
		for _, i := range num {
			buf <- i
		}
		close(buf)
	}()

	for v := range buf {
		fmt.Println(v)
	}
	fmt.Println(time.Since(t).Nanoseconds())

}
