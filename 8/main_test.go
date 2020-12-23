package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test1_AdjGreatestProduct(t *testing.T) {
	res := AdjGreatestProduct(4)
	assert.Equal(t, "9989", res.String(), "they should be equal")
	t.Logf("\nResult: %s\n", res)
}

func Test2_AdjGreatestProduct(t *testing.T) {
	res := AdjGreatestProduct(13)
	assert.Equal(t, "5576689664895", res.String(), "they should be equal")
	t.Logf("\nResult: %s\n", res)
}

var result *AdjResult

func Benchmark1_Prime_Nst(b *testing.B) {
	var r *AdjResult
	for n := 0; n < b.N; n++ {
		r = AdjGreatestProduct(4)
	}
	result = r
}

