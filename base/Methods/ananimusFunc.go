package main

import "fmt"

func main() {
	f:= increment()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(increment()())
	fmt.Println(increment()())
	fmt.Println(increment()())
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
