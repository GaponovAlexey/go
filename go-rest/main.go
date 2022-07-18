package main

import (
	"fmt"
)

func main() {
	s := NewSt()
	s.Name = "ddsd"
	s.Met1()

}

type st1 struct {
	Name string
	Age  int
}

type St2 struct {
	st1
	Title string
}

func (s St2) Met1() {
	fmt.Println(s.Name)

}
func (s St2) Met2() {
	fmt.Println(s.Name)
}

type int1 interface {
	Met1()
	Met2()
}

func NewSt() *St2 {
	return &St2{}
}
