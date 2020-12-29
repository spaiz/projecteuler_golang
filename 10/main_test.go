package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test1_SumOfPrimesBelow1(t *testing.T) {
	res := SumOfPrimesBelow1(10)
	assert.Equal(t, 17, res, "they should be equal")
}

func Test1_SpecialPythagoreanTriplet2(t *testing.T) {
	res := SumOfPrimesBelow1(2 * 1000000)
	assert.Equal(t, 142913828922, res, "they should be equal")
}

var result int

func Benchmark1_AdjGreatestProduct(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = SumOfPrimesBelow1(2*1000000)
	}
	result = r
}