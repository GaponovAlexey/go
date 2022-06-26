package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	bob := Cats{"Bob", 7, 0.92}
	fmt.Println(bob.happiness)
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Clock())
	fmt.Println(rand.Intn(10))
	

}

type Cats struct {
	name      string
	age       int
	happiness float64
}
