package main

import "fmt"

type R interface {
	F() string
}

type T struct {
	ST int
}
type TT struct {
	TT string
}

func (s *T) F() string {
	return fmt.Sprintf("%#v", s)
}
func (s *TT) F() string {
	return fmt.Sprintf("%#v", s)
}

func main() {
	var r R    //interface
	r = &TT{}      // присвоение
	r = &T{22}

	TypeAssertion(r)
}

func TypeAssertion(r R) {
	if t, ok := r.(*T); ok {
		fmt.Println(t.F())
	}}
