package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Fib1(t *testing.T) {
	res := Fib1(1, 2, 4000000, 0)
	assert.Equal(t, res, 4613730, "they should be equal")
}


func Test_Fib2(t *testing.T) {
	res := Fib2(1, 2, 4000000)
	assert.Equal(t, res, 4613730, "they should be equal")
}


func Test_Fib3(t *testing.T) {
	res := Fib3(1, 2, 4000000)
	assert.Equal(t, res, 4613730, "they should be equal")
}
