package main

import (
	"fmt"
	"math"
)

type Ver struct {
	X, t, y, u, i int
}

func (s Ver) add() int {
	return 22 + s.t
}
func (hera Ver) ass(s int) int {
	return hera.X + s + s
}

// type MF float64
type MF int

func (f MF) ADS() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	//methods
	data := Ver{22, 23, 24, 25, 26}
	fmt.Println(data.add())
	fmt.Println(data.ass(22))
	fmt.Println(data)
	ff := MF(-2)
	fmt.Println(ff.ADS())
	v := Vertex{3, 4}
	v.Scale(1)
	fmt.Println(v.Abs())

}
