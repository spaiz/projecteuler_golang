package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test1_SpecialPythagoreanTriplet1(t *testing.T) {
	res := SpecialPythagoreanTriplet1(1000)
	assert.Equal(t, 31875000, res.Product(), "they should be equal")
	t.Logf("\nResult: %d^2 + %d^2 + %d^2 = %d\n", res.a, res.b, res.c, res.Product())
}

func Test1_SpecialPythagoreanTriplet2(t *testing.T) {
	res := SpecialPythagoreanTriplet2(1000)
	assert.Equal(t, 31875000, res.Product(), "they should be equal")
	t.Logf("\nResult: %d^2 + %d^2 + %d^2 = %d\n", res.a, res.b, res.c, res.Product())
}

var result int

func Benchmark1_AdjGreatestProduct(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = SpecialPythagoreanTriplet1(1000).Product()
	}
	result = r
}

func Benchmark1_SpecialPythagoreanTriplet2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = SpecialPythagoreanTriplet2(1000).Product()
	}
	result = r
}