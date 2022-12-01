package main

import (
	"fmt"
	"gaponovalexey/mod/data"

)

func main() {
	newdata := data.Data
	fmt.Println(newdata)
	data.DataPackage("ivan")

}
