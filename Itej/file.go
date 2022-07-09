package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// create
	data:= []byte("text to file")
	err := ioutil.WriteFile("Itej/text.txt", data, 0600 )
// read
	file_data, err := ioutil.ReadFile("Itej/text.txt")
	if err != nil {
		fmt.Println("ne mogu pro4itat")
	}
	fmt.Println(string(file_data))

	os.Remove("Itej/text.txt")
}
