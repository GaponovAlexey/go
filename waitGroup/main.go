package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doSomething(a int) {
	fmt.Println("My job is", a)
	return
}
func main() {
	var ch = make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			for {
				a, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				fmt.Println(a)
			}
		}()
	}
	for i := 0; i < 50; i++ {
		ch <- i
	}

	go func() {
		close(ch)
		wg.Wait()
	}()
}
