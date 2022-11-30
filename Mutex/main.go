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
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	
	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func() {
			m.Lock()
			counter++
			m.Unlock()
			defer wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}