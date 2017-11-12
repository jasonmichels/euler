package palindrome

import (
	"math"
	"strconv"
)

// LargestPalindrome Find the largest palindrome made from the product of two 3-digit numbers
func LargestPalindrome() int {
	var largest int

	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			result := i * j
			if IsPalindrome(strconv.Itoa(result)) {
				if result > largest {
					largest = result
					break
				}
			}
		}
	}

	return largest
}

// IsPalindrome Check if the string is a palindrome
func IsPalindrome(s string) bool {
	// Strategy is to split the string in half, reverse the last half, then compare the strings

	if len(s) <= 1 {
		return false
	}

	middle := int(math.Floor(float64(len(s) / 2)))
	begin := s[:middle]

	var end string
	if len(s)%2 == 0 {
		end = s[middle:]
	} else {
		end = s[middle+1:]
	}

	reversed := make([]byte, len(end))

	j := 0
	for i := len(end) - 1; i >= 0; i-- {
		reversed[j] = end[i]
		j++
	}

	return begin == string(reversed[:])
}
