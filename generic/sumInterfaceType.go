package main

import "fmt"

type Number interface {
	~int64 | float64
}
type CustInt int64

func (ci CustInt) IsPositive() bool {
	return ci > 0
}

type Numbers[T Number] []T

func main() {
	unionInterfaceAndType()
}

func unionInterfaceAndType() {
	var ints Numbers[int64]
	ints = append(ints, []int64{1, 2, 3, 4, 5}...)

	floats := Numbers[float64]{1.0, 2, 5, 6, 1.1}
	fmt.Println(sumUnionInterface(ints))
	fmt.Println(sumUnionInterface(floats))
}

func sumUnionInterface[V Number](num []V) V {
	var sum V
	for _, v := range num {
		sum += v
	}
	return sum
}
