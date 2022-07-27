package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var c int64
	
	atomic.StoreInt64(&c, 6)
	fmt.Println(atomic.LoadInt64(&c))
	
	fmt.Println(atomic.SwapInt64(&c, 10))
	fmt.Println(atomic.LoadInt64(&c))
	

}
