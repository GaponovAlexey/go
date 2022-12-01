package fib_test

import (
	"gaponovalexey/mod/fib"
	"testing"

	"github.com/stretchr/testify/assert"

)

func TestFib(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "zero",
			n:    0,
			want: 0,
		},
		{
			name: "one",
			n:    1,
			want: 1,
		},
		{
			name: "two",
			n:    2,
			want: 1,
		},
		{
			name: "one",
			n:    3,
			want: 2,
		},

	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, fib.Fib(tc.n))
		})
	}
}
