package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.UTC().Day()
	fmt.Println(time)
	
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
