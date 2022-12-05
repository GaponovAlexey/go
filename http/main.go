package main

import (
	"fmt"
	"io"
	"net/http"

)

func main() {
	c := &http.Client{}
	g, _ := c.Get("https://jsonplaceholder.typicode.com/todos")
	bo, _ := io.ReadAll(g.Body)
	fmt.Println(bo)
	defer g.Body.Close()

	// fmt.Println(c.client.Get("https://jsonplaceholder.typicode.com/todos"))
}
