package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1_Prime_Nst(t *testing.T) {
	res := Prime_Nst(10001)
	assert.Equal(t, res, 104743, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

var result int

func Benchmark1_Prime_Nst(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Prime_Nst(100)
	}
	result = r
}