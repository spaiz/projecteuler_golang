package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1_SmallestMultiple(t *testing.T) {
	res := SmallestMultiple(1,10)
	assert.Equal(t, 2520, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test2_SmallestMultiple(t *testing.T) {
	res := SmallestMultiple(1,20)
	assert.Equal(t, 232792560, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test1_SmallestMultiple2(t *testing.T) {
	res := SmallestMultiple2(1,20)
	assert.Equal(t, 232792560, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test1_SmallestMultiple3(t *testing.T) {
	res := SmallestMultiple3(1,20)
	assert.Equal(t, 232792560, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

func Test1_SmallestMultiple4(t *testing.T) {
	res := SmallestMultiple4(1,20)
	assert.Equal(t, 232792560, res, "they should be equal")
	t.Logf("\nResult: %d\n", res)
}

var result int

func Benchmark1_SmallestMultiple(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = SmallestMultiple(1, 20)
	}
	result = r
}

func Benchmark1_SmallestMultiple2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = SmallestMultiple2(1, 20)
	}
	result = r
}

func Benchmark1_SmallestMultiple3(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = SmallestMultiple3(1, 40)
	}
	result = r
}

func Benchmark1_SmallestMultiple4(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = SmallestMultiple4(1, 40)
	}
	result = r
}