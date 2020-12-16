package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1_Diff1(t *testing.T) {
	res := Diff1(10)
	assert.Equal(t, 2640, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test2_Diff1(t *testing.T) {
	res := Diff1(100)
	assert.Equal(t, 25164150, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test1_Diff2(t *testing.T) {
	res := Diff2(10)
	assert.Equal(t, 2640, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test2_Diff2(t *testing.T) {
	res := Diff2(100)
	assert.Equal(t, 25164150, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

var result int

func Benchmark1_Diff1(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Diff1(100)
	}
	result = r
}

func Benchmark1_Diff2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = Diff2(100)
	}
	result = r
}
