package smallestmultiple

import (
	"github.com/jasonmichels/euler/lcm"
)

// Naive Find smallest multiple of numbers 1-20, naive solution
func Naive() int {
	//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
	// From 1 to 10 it is 2520

	checks := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var smallest int

	for i := 2520; i < 1000000000; i++ {
		passed := true
		for _, v := range checks {
			if i%v != 0 {
				passed = false
				break
			}
		}
		if passed {
			smallest = i
			break
		}
	}

	return smallest
}

// NaiveBetter Has much better performance since removed alot of numbers to check, but can do better
func NaiveBetter() int {
	checks := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var smallest int

	for i := 2520; i < 1000000000; i = i + 10 {
		passed := true
		for _, v := range checks {
			if i%v != 0 {
				passed = false
				break
			}
		}
		if passed {
			smallest = i
			break
		}
	}

	return smallest
}

// WithLCM With Least common multiple function, and GCD function I added earlier
func WithLCM() int {
	var result = 11
	for i := 11; i < 20; i++ {
		result = lcm.LCM(result, i+1)
	}
	return result
}
