package main

import (
	"fmt"
	"github.com/GaponovAlexey/mysum"
)

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
	fmt.Println(g.add(name))
}
func NewF() Fantazer {
	return &fanta{}
}
func main() {
	// name := "zer"
	// f := NewF()
	data := mysum.Sum(2,5)
	fmt.Println(data)
	
	// seyHello(&fanta{}, name)
	// seyHello(&vodka{}, name)

}
