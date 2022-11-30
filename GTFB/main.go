package main

import (
	"fmt"
	"log"
)

func main() {
	country := "london"

	c, err := callName(country)

	log.Println(c)
	log.Println(err)

}

func callName(c string) (string, error) {
	c = c + " " + "counrty"
	// newc := append(c, "country")
	return c, fmt.Errorf("error %v", c)
}
