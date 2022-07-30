package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	workerPool()
}

func workerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	f, s := make(chan int, 5), make(chan int, 5)
	wg := &sync.WaitGroup{}

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, f, s)
		}()
	}

	go func() {
		for i := 0; i <= 20; i++ {
			f <- i
		}
		close(f)
	}()

	go func() {
		wg.Wait()
		close(s)
	}()
	for s := range f {
		fmt.Println("tp", s)
	}
	for i := range s {
		fmt.Println("pns", i)
	}

}

func worker(ctx context.Context, f <-chan int, s chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case f, ok := <-f:
			if !ok {
				return
			}
			s <- f + f
		}
	}
}
