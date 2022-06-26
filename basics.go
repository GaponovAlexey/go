package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func main() {
	//type
	bob := Cats{"Bob", 7, 0.92}
	fmt.Println(bob.happiness)
	//time
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Clock())
	//math - random Determinism
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Sqrt(1))
	fmt.Printf("Now you have sqrt %g problems.\n", math.Sqrt(2))
	fmt.Println(math.Pi)
	//swap
	fmt.Println(swap("hello", "go"))
	//split
	fmt.Println(split(17))
	//create in out
	i, j := 1, 2
	fmt.Println(i, j, c, pytohon, java)
	//Printf %T %v 
	fmt.Printf("type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	//type %q string
	b, s := 22, "dsad"
	fmt.Printf("%v %q\n", b, s)
	
	x, y := 3, 5
	f := math.Sqrt(float64(x*y + y*y))
	z := (f)
	fmt.Println(x, y, z);
	
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
