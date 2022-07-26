package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	workerPool()
}

func workerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(context.Background(), time.Microsecond*20)
	defer cancel()

	wg := &sync.WaitGroup{}

	tp, pns := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, tp, pns)
		}()
	}

	go func() {
		for i := 0; i <= 1000; i++ {
			tp <- i
		}
		close(tp)
	}()

	go func() {
		wg.Wait()
		close(pns)
	}()

	var counter int
	for res := range pns {
		counter++
		fmt.Println(res)
	}
	fmt.Println(counter)

}

func worker(ctx context.Context, toP <-chan int, proc chan<- int) {
	for {
		select {
		case <-ctx.Done(): // дожидается завершения контекста
			return
		case v, ok := <-toP:
			if !ok {
				return
			}
			time.Sleep(time.Microsecond)
			proc <- v * v
		}
	}
}
