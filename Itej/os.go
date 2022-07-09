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
	data := "you sss"      //text
	file.WriteString(data) // write

	// file_data, _ := os.ReadFile(file.Name()) // отсылка на созданый файл
	file_data, _ := os.ReadFile("text.txt") // отсылка на созданый файл

	fmt.Println(string(file_data))

	//defer
	defer file.Close()
}
