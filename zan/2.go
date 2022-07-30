package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

func main() {
	t := time.Now()
	wg := sync.WaitGroup{}
	key := "fest"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			c.Inc(key)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.Value(key))
	fmt.Println(time.Since(t).Seconds())
	
}
