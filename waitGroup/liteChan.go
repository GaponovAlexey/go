package main

import (
	"fmt"
	"sync"
)

func main() {
	chanAsMutex()
}


func chanAsMutex() {
	var counter int
	mutexChan := make(chan struct{}, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mutexChan <- struct{}{}

			counter++

			<-mutexChan
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}