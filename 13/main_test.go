package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test1_GreatestProduct1(t *testing.T) {
	res := FirstDigitsOfSum()
	assert.Equal(t, "5537376230", res.GetDigits(10), "they should be equal")
}

var result *BigNum

func Benchmark1_GreatestProduct1(b *testing.B) {
	var r *BigNum
	for n := 0; n < b.N; n++ {
		r = FirstDigitsOfSum()
	}
	result = r
}
