package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test1_GreatestProduct1(t *testing.T) {
	res := FirstDigitsOfSum()
	assert.Equal(t, "5537376230", res.GetDigits(10), "they should be equal")
}

func Test1_GreatestProduct2(t *testing.T) {
	res := FirstDigitsOfSum2(10)
	assert.Equal(t, "5537376230", res, "they should be equal")
}

var result *BigNum

func Benchmark1_GreatestProduct1(b *testing.B) {
	var r *BigNum
	for n := 0; n < b.N; n++ {
		r = FirstDigitsOfSum()
	}
	result = r
}

var result2 string
func Benchmark1_GreatestProduct2(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = FirstDigitsOfSum2(10)
	}
	result2 = r
}
