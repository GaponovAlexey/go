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
//func
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
//
// builder
func (b computerBuilder) Builder() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

func main() {
	cb := NewCB()
	da := cb.
	fmt.Println(da)

	net := NewCF()
	net.RAM(22)
	n := net.Builder()
	fmt.Println(n)

}

type officeComputerBuilder struct {
	computerBuilder
}

func NewCF() ComputerBuilder {
	return &officeComputerBuilder{}
}

func (b *officeComputerBuilder) Builder() Computer {
	if b.cpu == "" {
		b.CPU("Pentium")
	}
	if b.ram == 0 {
		b.RAM(2)
	}
	if b.mb == "" {
		b.MB("gigabate")
	}
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}
