package main

func main() {
	num := []int{5, 5, 6, 7}
	c := make(chan int)
	go func() {
		for _, i := range num {
			c <- i
		}
		close(c)
	}()
}
