package main

import "fmt"

type AllInter interface {
	first
	second
}
type first interface {
	add()
}
type second interface {
	add()
}

type allS struct {
	first
	second
}

type r struct{}
type a struct{}

func (v r) add() {
	fmt.Println("Ru")
}
func (v a) add() {
	fmt.Println("Am")
}

func NewI() AllInter {
	return r{}
}

func main() {
	s := NewI()
	s.add()
}
