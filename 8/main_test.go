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

func Test1_AdjGreatestSkipZeroMultiply(t *testing.T) {
	res := AdjGreatestSkipZeroMultiply(4)
	assert.Equal(t, "9989", res.String(), "they should be equal")
	t.Logf("\nResult: %s\n", res)
}

func Test2_Test1_AdjGreatestSkipZeroMultiply(t *testing.T) {
	res := AdjGreatestSkipZeroMultiply(13)
	assert.Equal(t, "5576689664895", res.String(), "they should be equal")
	t.Logf("\nResult: %s\n", res)
}

func Test1_AdjGreatestProductSkipZero(t *testing.T) {
	res := AdjGreatestProductSkipZero(4)
	assert.Equal(t, "9989", res.String(), "they should be equal")
	t.Logf("\nResult: %s\n", res)
}

func Test2_AdjGreatestProductSkipZero(t *testing.T) {
	res := AdjGreatestProductSkipZero(13)
	assert.Equal(t, "5576689664895", res.String(), "they should be equal")
	t.Logf("\nResult: %s\n", res)
}

func Test1_AdjGreatestProductSlidingWindow(t *testing.T) {
	res := AdjGreatestProductSlidingWindow(4)
	assert.Equal(t, "9989", res.String(), "they should be equal")
	t.Logf("\nResult: %s\n", res)
}

func Test2_AdjGreatestProductSlidingWindow(t *testing.T) {
	res := AdjGreatestProductSlidingWindow(13)
	assert.Equal(t, "5576689664895", res.String(), "they should be equal")
	t.Logf("\nResult: %s\n", res)
}

var result *AdjResult

func Benchmark1_AdjGreatestProduct(b *testing.B) {
	var r *AdjResult
	for n := 0; n < b.N; n++ {
		r = AdjGreatestProduct(13)
	}
	result = r
}

func Benchmark1_AdjGreatestSkipZeroMultiply(b *testing.B) {
	var r *AdjResult
	for n := 0; n < b.N; n++ {
		r = AdjGreatestSkipZeroMultiply(13)
	}
	result = r
}

func Benchmark1_AdjGreatestProductSkipZero(b *testing.B) {
	var r *AdjResult
	for n := 0; n < b.N; n++ {
		r = AdjGreatestProductSkipZero(13)
	}
	result = r
}

func Benchmark1_AdjGreatestProductSlidingWindow(b *testing.B) {
	var r *AdjResult
	for n := 0; n < b.N; n++ {
		r = AdjGreatestProductSlidingWindow(13)
	}
	result = r
}
