package smallestmultiple

import (
	"testing"
)

func TestNaiveIsCorrect(t *testing.T) {
	if result := Naive(); result != 232792560 {
		t.Errorf("Did not expect to receive: %v", result)
	}
}

func BenchmarkNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Naive()
	}
}

func TestNaiveBetterIsCorrect(t *testing.T) {
	if result := NaiveBetter(); result != 232792560 {
		t.Errorf("Did not expect to receive: %v", result)
	}
}

func BenchmarkNaiveBetter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NaiveBetter()
	}
}

func TestWithLCMIsCorrect(t *testing.T) {
	if result := WithLCM(); result != 232792560 {
		t.Errorf("Did not expect to receive: %v", result)
	}
}

func BenchmarkWithLCM(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WithLCM()
	}
}
