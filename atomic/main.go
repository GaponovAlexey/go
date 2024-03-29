package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	add()
	addAtomic()

}

func add() {
	s := time.Now()
	var (
		c  int64
		wg sync.WaitGroup
		m  sync.Mutex
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			m.Lock()
			c++
			c--
			c++
			m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(c)

	fmt.Println(time.Since(s).Seconds())
}

var (
	c  int64
	wg sync.WaitGroup
)
func addAtomic() {
	s := time.Now()
	wg.Add(1000)
	fmt.Println("Atomic:",atomic.LoadInt64(&c))
	

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&c, 1)
			// atomic.AddInt64(&c, -1)
			// atomic.AddInt64(&c, 1)
		}()
	}
	wg.Wait()
	fmt.Println(c)

	fmt.Println(time.Since(s).Seconds())
}
