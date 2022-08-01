package main

import "fmt"
type all interface {
	animal
	bird
}

type animal interface {
	walker
	runner
}
type bird interface {
	walker
	flier()
}
type walker interface {
	walk()
}
type runner interface {
	run()
}

//struct
type cat struct{}

//func
func (c *cat) walk() {
	fmt.Println("cat is walking")
}
func (c *cat) run() {
	fmt.Println("cat is walking")
}
func (c *cat) flier() {
	fmt.Println("bird is flier")

}
func main() {
	var c all = &cat{}
	c.walk()
	c.run()
	c.walk()	
}
