package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Iterative(t *testing.T) {
	n := 10
	res := FindIter(n)
	assert.Equal(t, res, 23, "they should be equal")
}

func Test_Recursive(t *testing.T) {
	n := 10
	res := FindRec(n - 1)
	assert.Equal(t, res, 23, "they should be equal")
}