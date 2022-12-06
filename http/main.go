package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

)

func main() {
	c := &http.Client{}
	g, _ := c.Get("https://jsonplaceholder.typicode.com/todos")
	bo, _ := io.ReadAll(g.Body)
	fmt.Println(string(bo))
	defer g.Body.Close()


	//save in folder
	fPath, nameFile := "storage", "file.json"
	joi := filepath.Join(fPath, nameFile)

	_ = os.MkdirAll(fPath, 0774)
	_ = os.WriteFile(joi, []byte(bo), 0774)

	file, _ := os.ReadFile(fPath)
	os.Stdout.Write(file)
	fmt.Println(file)

}

