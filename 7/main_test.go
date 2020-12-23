package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1_Prime_Nst(t *testing.T) {
	res := PrimeNst(10001)
	assert.Equal(t, res, 104743, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test1_Prime_Nst_Sieve(t *testing.T) {
	res := PrimeSeieve1(10001)
	assert.Equal(t, 104743, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test1_Prime_Nst_Sieve2(t *testing.T) {
	res := PrimeSeieve2(10001)
	assert.Equal(t, 104743, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

var result int

func Benchmark1_Prime_Nst(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = PrimeNst(100)
	}
	result = r
}

func Benchmark1_Prime_Sieve(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = PrimeNst(100)
	}
	result = r
}

func Benchmark1_Prime_Sieve2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = PrimeSeieve2(100)
	}
	result = r
}