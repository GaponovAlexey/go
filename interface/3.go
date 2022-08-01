package main

import "fmt"

type animal interface { //корневой
	walker //1
	runner //2
}
type bird interface { //корневой
	walker //1
	flier  //3
}

type walker interface {
	walk() //1
}
type runner interface {
	run() //2
}
type flier interface {
	fly() //3
}

//struct
type cat struct{}
type eagle struct{}

//func
func (c *cat) walk() {
	fmt.Println("cat is walking")
}
func (c *cat) run() {
	fmt.Println("cat is walking")
}

//eagle
func (e *eagle) walk() {
	fmt.Println("eagle is walk")
}
func (e *eagle) fly() {
	fmt.Println("eagle is fly")
}
func walk(w walker) {
	w.walk()
}

func main() {
	var c animal = &cat{}
	var e bird = &eagle{}
	walk(c)
	walk(e)
}
