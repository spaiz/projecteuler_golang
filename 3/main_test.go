package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var result int

func Test1_LPF1(t *testing.T) {
	res := LPF1(13195)
	assert.Equal(t, 29, res, "they should be equal")
}

func Test2_LPF1(t *testing.T) {
	res := LPF1(181955)
	assert.Equal(t, 241, res, "they should be equal")
}

func Test3_LPF1(t *testing.T) {
	res := LPF1(600851475143)
	assert.Equal(t, 6857, res, "they should be equal")
}

func Test1_LPF2(t *testing.T) {
	res := LPF2(600851475143)
	assert.Equal(t, 6857, res, "they should be equal")
}

func Test1_LPF3(t *testing.T) {
	res := LPF3(600851475143)
	assert.Equal(t, 6857, res, "they should be equal")
}

func Benchmark3_LPF1(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = LPF1(n)
	}
	result = r
}

func Benchmark1_LPF2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = LPF2(n)
	}
	result = r
}

func Benchmark1_LPF3(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		r = LPF2(n)
	}
	result = r
}