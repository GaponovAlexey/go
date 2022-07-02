package main

import (
	"fmt"
)

func main() {
	data := []string{
		"mes_1",
		"mes_2",
		"mes_3",
		"mes_4",
	}
	fmt.Println(data)
	c := 0
	for {
		c++
		fmt.Println(c)
		if c == 100 {
			break
		}
	}

}
