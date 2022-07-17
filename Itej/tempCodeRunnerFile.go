func say(great string, ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
