package main

import "fmt"



type ComputerBuilder interface {
	CPU(val string) ComputerBuilder
	RAM(val int) ComputerBuilder
	MB(val string) ComputerBuilder

}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func NewCB() ComputerBuilder {
	return &computerBuilder{}
}

func (b computerBuilder) CPU(val string) ComputerBuilder {
	b.cpu = val
	return b
}
func (b computerBuilder) RAM(val int) ComputerBuilder {
	b.ram = val
	return b
}
func (b computerBuilder) MB(val string) ComputerBuilder {
	b.mb = val
	return b
}


func main() {
	cb := NewCB()
	da := cb.CPU("cor i9").RAM(64).MB("gigabate")
	fmt.Println(da)

}

