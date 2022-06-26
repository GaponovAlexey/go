package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func main() {
	bob := Cats{"Bob", 7, 0.92}
	fmt.Println(bob.happiness)
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Clock())
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Sqrt(1))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(2))
	fmt.Println(math.Pi)
	fmt.Println(swap("hello", "go"))
	fmt.Println(split(17))
	i, j := 1, 2
	fmt.Println(i, j, c, pytohon, java)
	fmt.Printf("type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	b, s := 22, "dsad"
	fmt.Printf("%v %q\n", b, s)

}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// create
const c, pytohon, java = true, false, "string!"

// swap
func swap(x string, y string) (any, any) {
	return x, y
}

// split
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// type
type Cats struct {
	name      string
	age       int
	happiness float64
}
