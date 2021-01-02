package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1_LongestCollatzSequence(t *testing.T) {
	res := LongestCollatzSequence(1000000)
	assert.Equal(t, 525, res.ChainSize(), "they should be equal")
	assert.Equal(t, 837799, res.GetStart(), "they should be equal")

	for _, v := range res.Chain() {
		t.Logf("%d ->", v)
	}
}

func Test1_LongestCollatzSequenceMemoize(t *testing.T) {
	res := LongestCollatzSequenceMemoize(1000000)
	assert.Equal(t, 525, res.ChainSize(), "they should be equal")
	assert.Equal(t, 837799, res.GetStart(), "they should be equal")

	for _, v := range res.Chain() {
		t.Logf("%d ->", v)
	}
}

var result *Result

func Benchmark1_LongestCollatzSequence(b *testing.B) {
	var r *Result
	for n := 0; n < b.N; n++ {
		r = LongestCollatzSequence(1000000)
	}
	result = r
}

func Benchmark1_LongestCollatzSequenceMemoize(b *testing.B) {
	var r *Result
	for n := 0; n < b.N; n++ {
		r = LongestCollatzSequenceMemoize(1000000)
	}
	result = r
}
