package problem06

import (
	"testing"
)

func TestSumSquaresTen(t *testing.T) {
	if result := SumSquares(10); result != 385 {
		t.Errorf("Expected: 385, Received: %v", result)
	}
}

func BenchmarkTestSumSquaresTen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumSquares(10)
	}
}

func TestSquareSumsTen(t *testing.T) {
	if result := SquareSums(10); result != 3025 {
		t.Errorf("Expected: 3025, Received: %v", result)
	}
}

func BenchmarkSquareSumsTen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SquareSums(10)
	}
}

func BenchmarkTestSumSquaresHundred(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumSquares(100)
	}
}

func BenchmarkSquareSumsHundred(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SquareSums(100)
	}
}

func TestSumSquareDiff(t *testing.T) {
	if result := SumSquareDiff(); result != 25164150 {
		t.Errorf("Expected: 25164150, Received: %v", result)
	}
}

func BenchmarkSumSquareDiff(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumSquareDiff()
	}
}

func TestSumSquareDiffChannel(t *testing.T) {
	if result := SumSquareDiffChannel(); result != 25164150 {
		t.Errorf("Expected: 25164150, Received: %v", result)
	}
}

func BenchmarkSumSquareDiffChannel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SumSquareDiffChannel()
	}
}
