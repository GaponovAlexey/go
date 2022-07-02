package main

import (
	"fmt"
)

type st struct {
	name string
}

func main() {
	d := st{}
	fmt.Printf("%+v",d)

}
