package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	//say
	go say("hello go", &ch)

	fmt.Println(<-ch)

}
func say(great string, ch *chan int) {
	data := 20
	*ch <- data
	close(*ch)
}
