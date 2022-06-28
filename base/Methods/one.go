package main

import (
	"fmt"
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
type MF float64

func (f MF) ADS() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	//methods
	data := Ver{22, 23, 24, 25, 26}
	fmt.Println(data.add())
	fmt.Println(data.ass(22))
	fmt.Println(data)

}
