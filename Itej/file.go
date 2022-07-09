package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("file")
	file_data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("ne mogu pro4itat")
	}
	fmt.Println(string(file_data))

}
