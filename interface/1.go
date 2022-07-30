package main

import "fmt"

type Fantazer interface {
	add(string) string
}

type fanta struct{}
type vodka struct{}

func (v vodka) add(name string) string {
	return fmt.Sprintln("Vodka:", name)
}
func (f fanta) add(name string) string {
	return fmt.Sprintln("Fanta:", name)
}

func seyHello(g Fantazer, name string) {
	fmt.Println(g.add(name))// оборачивает add()

}
func main() {
	name := "zer"
	seyHello(&fanta{}, name)
	seyHello(&vodka{}, name)

}
