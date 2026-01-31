package tests

import (
	"math/rand"
	"testing"
)

func GenerateAndSumRandomNumSlice(size int) int {
	randInts := make([]int, size)
	sum := 0

	for i := range size {
		randInts[i] = rand.Intn(100)
	}

	for _, val := range randInts {
		sum += val
	}

	return sum
}

func BenchmarkGenAndSumRandInts(b *testing.B) {
	for b.Loop() {
		GenerateAndSumRandomNumSlice(100)
	}
}