package main

import (
	"fmt"
	"log"
)

func main() {
	var userName string
	var books uint8

	books = 3
	books = 255
	// fmt.Scan(&userName)
	fmt.Scan(books, userName)
	log.Println(&userName)

}
