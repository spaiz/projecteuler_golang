package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestPalindrome1(t *testing.T) {
	res, i, y := LargestPalindrome1(2)
	assert.Equal(t, 9009, res, "they should be equal")
	t.Logf("\nResult: %d*%d\n", i, y)
}

func TestLargestPalindrome2(t *testing.T) {
	res, i, y := LargestPalindrome1(3)
	assert.Equal(t, 906609, res, "they should be equal")
	t.Logf("\nResult: %d*%d\n", i, y)
}