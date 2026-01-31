package tests

import "testing"

func add(a, b int) int {
	return a + b
}

func BenchmarkAddSmall(b *testing.B) {
	for b.Loop() {
		Add(2, 3)
	}
}

func BenchmarkAddMedium(b *testing.B) {
	for b.Loop() {
		Add(2000, 3000)
	}
}

func BenchmarkAddLarge(b *testing.B) {
	for b.Loop() {
		Add(200_000_000, 300_000_000)
	}
}
