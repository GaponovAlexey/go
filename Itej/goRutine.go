package main

import "fmt"

type Num struct {
	n1 int
	n2 int
}
type NumberIterfece interface {
	Sum() int
	Min() int
}

func (n Num) Sum() int {
	return n.n1 + n.n2
}
func (n Num) Min() int {
	return n.n1 - n.n2
}

func main() {
	var i NumberIterfece
	numbers := Num{5, 2}
	i = numbers
	fmt.Println(i.Sum())
	fmt.Println(i.Min())

}
