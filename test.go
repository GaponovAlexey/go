package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// ch2 := make(chan int)
	go rut1(ch)

	for {
		fmt.Println(<-ch)
		break
	}

	time.Sleep(time.Second * 2)
}

func rut1(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		
		// time.Sleep(time.Microsecond)
		ch <- i
	}
	defer close(ch)
}
