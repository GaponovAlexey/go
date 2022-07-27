package main

import (
	"fmt"
	"time"
)

func main() {
	chanAsPromise()

}


func fr(num int) <-chan string {
	rChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		rChan <- fmt.Sprintf("res num %v", num)
	}()
	return rChan
}
func chanAsPromise() {
	f := fr(1)
	f2 := fr(2)
	fmt.Println("not blocking")
	fmt.Println(<-f, <-f2)
}
