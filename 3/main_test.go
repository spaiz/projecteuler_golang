package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

