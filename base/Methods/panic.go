package main

import (
	"fmt"
)

func main() {
	panic()

	fmt.Println("main()")

}

func panic() {
	fmt.Println("print defer")
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
