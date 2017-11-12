package gcd

import (
	"testing"
)

func TestGCDFor24And54(t *testing.T) {
	if result := GCD(24, 54); result != 6 {
		t.Errorf("Was expecting GCD(24,54) to be 6, not: %v", result)
	}
}

func BenchmarkGCDFor24And54(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GCD(24, 54)
	}
}

func TestGCDFor24And54Recursive(t *testing.T) {
	if result := Recursive(24, 54); result != 6 {
		t.Errorf("Was expecting GCD(24,54) to be 6, not: %v", result)
	}
}

func BenchmarkGCDFor24And54Recursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Recursive(24, 54)
	}
}
