package main

import "fmt"

func main() {
	bob:= Cats {"Bob", 7, 0.92}
	fmt.Println(bob.age);
}
type Cats struct {
	name string
	age int
	happiness float64
}