package main

import (
	"fmt"
)

func main() {
	fmt.Printf("ex1:%T ex2:%T", example(), example2())

}

type Example struct {
	Value string
}
type MyInterface interface{}

func example() MyInterface {
	var e *Example
	return e
}

func example2() MyInterface {
	return nil
}
