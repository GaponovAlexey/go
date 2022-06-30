package main

import (
	"errors"
	"fmt"
)

func main() {
	mass := [3]string{"1", "2", "3"}
	print(mass)

}

func print(mass [3]string) error {
	if len(mass) == 0 {
		return errors.New("Nothing")
	}
	mass[1] = "5"
	fmt.Println(mass)
	return nil

}
