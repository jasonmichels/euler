package primefactor

import (
	"testing"
)

func TestCanGetPrimeFactors(t *testing.T) {
	expected := []int{5, 7, 13, 29}

	result := PrimeFactors(13195)
	if len(result) != 4 {
		t.Errorf("Unexpected result of %v", result)
	}

	if result[0] != expected[0] && result[1] != expected[1] && result[2] != expected[2] && result[3] != expected[3] {
		t.Errorf("Unexpected result of %v", result)
	}
}

func BenchmarkCanGetPrimeFactors(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PrimeFactors(13195)
	}
}

func BenchmarkCanGetPrimeFactorsLargeNum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PrimeFactors(600851475143)
	}
}

func TestIfIsPrime(t *testing.T) {
	primes := []int{103, 107, 109, 113, 127}

	for i := 0; i < len(primes); i++ {
		if !IsPrime(primes[i]) {
			t.Error("This number is actually a prime")
		}
	}

}

func BenchmarkIfIsPrime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPrime(199)
	}
}
