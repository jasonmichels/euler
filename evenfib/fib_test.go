package evenfib

import (
	"testing"
)

func TestEvenFibCalculatesEvenFib(t *testing.T) {
	if sum := EvenFib(89); sum != 44 {
		t.Errorf("Unexpected result of %v, when I wanted 44", sum)
	}
}

func BenchmarkEvenFibUpTo89(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EvenFib(89)
	}
}

func TestEvenFibCalculatesUpToFourMillion(t *testing.T) {
	if sum := EvenFib(4000000); sum != 4613732 {
		t.Errorf("Max number: 4,000.  Unexpected result of %v, when I wanted 4613732", sum)
	}
}

func BenchmarkEvenFibUpToFourMillion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EvenFib(4000000)
	}
}

func BenchmarkEvenFibUpToLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EvenFib(7223372036854775807)
	}
}
