package primefactor

import (
	"math"
)

// PrimeFactors Calculate the prime factors for a number
func PrimeFactors(num int) []int {
	factors := []int{}

	// Get the number of 2s that divide num
	for num%2 == 0 {
		factors = append(factors, 2)
		num = num / 2
	}

	// num must be odd at this point
	for i := 3; i*i <= num; i = i + 2 {
		// while i divides num, append i and divide num
		for num%i == 0 {
			factors = append(factors, i)
			num = num / i
		}
	}

	// This condition is to handle the case when num is a prime number
	// greater than 2
	if num > 2 {
		factors = append(factors, num)
	}

	return factors
}

// IsPrime Determine if a number is prime
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	maxNum := int(math.Floor(math.Sqrt(float64(num))))

	for i := 3; i <= maxNum; i = i + 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
