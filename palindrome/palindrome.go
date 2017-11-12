package palindrome

import (
	"math"
	"strconv"
)

// LargestPalindrome Find the largest palindrome made from the product of two 3-digit numbers
func LargestPalindrome() int64 {
	var largest int64

	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			result := int64(i * j)

			if IsPalindrome(result) {
				if result > largest {
					largest = result
					break
				}
			}
		}
	}

	return largest
}

// IsPalindrome Check if the integer is a palindrome
func IsPalindrome(n int64) bool {
	// The strategry is to convert int to string
	// split the string in half, reverse the last half, then compare the strings

	s := strconv.FormatInt(n, 10)

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
