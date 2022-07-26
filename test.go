package main

import "fmt"

type Computer struct {
	CPU string
	RAM int
	MB  string
}

type ComputerBuilder interface {
	CPU(val string) ComputerBuilder
	RAM(val int) ComputerBuilder
	MB(val string) ComputerBuilder

	Builder() Computer
}

type computerBuilder struct {
	cpu string
	ram int
	mb  string
}

func NewCB() ComputerBuilder {
	return computerBuilder{}
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

func (b computerBuilder) Builder() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

func main() {
	cb := NewCB()
	da := cb.CPU("cor i9").RAM(64).MB("gigabate").Builder()
	fmt.Println(da)

}

type officeComputerBuilder struct {
	computerBuilder
}
