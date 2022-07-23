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
	t := &TT{} //type
	r = t      // присвоение

	tt := &T{22}
	r = tt

	TypeAssertion(tt)
	if r == nil {
		fmt.Printf("%#v %T", r, r)
	}

}

func TypeAssertion(r R) {
	if t, ok := r.(*T); ok {
		fmt.Println(t.F())
	}}
