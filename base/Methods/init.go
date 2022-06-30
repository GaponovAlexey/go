package main

import "fmt"

var msg string
func init() {
	msg = "This is init"
}

func main() {
	fmt.Println(msg)

}