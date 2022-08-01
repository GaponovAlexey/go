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
<<<<<<< HEAD:interface/2.go
	fmt.Println(g.add(name))
=======
	fmt.Println(g.add(name))// оборачивает add()

>>>>>>> 1b8c51674580b22e70ef90e30b6302073f949d98:interface/1.go
}

func main() {
	name := "zer"
	seyHello(&fanta{}, name)
	seyHello(&vodka{}, name)
}
