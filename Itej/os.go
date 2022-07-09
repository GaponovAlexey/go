package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt") // create
	if err != nil {
		fmt.Println("error file")
	}
	data := "you sss" //text
	file.WriteString(data) // write

	//defer
	defer file.Close()
}
