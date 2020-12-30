package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test1_GreatestProduct1(t *testing.T) {
	res := HighlyDivisibleTriangularNumber(5)
	assert.Equal(t, 28, res.number, "they should be equal")

	for _, v := range res.divisors {
		t.Logf("%d, ", v)
	}
}

func Test2_GreatestProduct1(t *testing.T) {
	t.Skip("Too slow")
	res := HighlyDivisibleTriangularNumber(500)
	assert.Equal(t, 76576500, res.number, "they should be equal")

	for _, v := range res.divisors {
		t.Logf("%d, ", v)
	}
}

func Test1_GreatestProduct2(t *testing.T) {
	res := HighlyDivisibleTriangularNumber2(500)
	assert.Equal(t, 76576500, res.number, "they should be equal")
}

var result *Result

func Benchmark1_GreatestProduct1(b *testing.B) {
	var r *Result
	for n := 0; n < b.N; n++ {
		r = HighlyDivisibleTriangularNumber(50)
	}
	result = r
}

func Benchmark1_GreatestProduct2(b *testing.B) {
	var r *Result
	for n := 0; n < b.N; n++ {
		r = HighlyDivisibleTriangularNumber2(50)
	}
	result = r
}