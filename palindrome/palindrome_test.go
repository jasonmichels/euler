package palindrome

import (
	"strconv"
	"testing"
)

func TestCanGetLargestPalindrome(t *testing.T) {
	if result := LargestPalindrome(); result != 906609 {
		t.Errorf("Unexpected result %v, was expecting: 906609", result)
	}
}

func BenchmarkCanGetLargestPalindrome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LargestPalindrome()
	}
}
func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome(strconv.Itoa(900009)) {
		t.Error("Expected 900009 to be a palindrome, but it wasnt")
	}
}

func TestIsPalindromeOdd(t *testing.T) {
	if !IsPalindrome(strconv.Itoa(9004009)) {
		t.Error("Expected 9004009 to be a palindrome, but it wasnt")
	}
}

func TestIsNotPalindrome(t *testing.T) {
	if IsPalindrome(strconv.Itoa(900089)) {
		t.Error("Expected 900089 to NOT be a palindrome")
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPalindrome(strconv.Itoa(900009))
	}
}

func BenchmarkIsPalindromeLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPalindrome(strconv.Itoa(9000000001000000009))
	}
}

func TestStringIsPalindrome(t *testing.T) {
	if !IsPalindrome("jasonnosaj") {
		t.Errorf("Expected jasonnosaj to be a palindrome")
	}
}

func TestStringIsNotPalindrome(t *testing.T) {
	if IsPalindrome("jasonisnotapalindrome") {
		t.Errorf("Expected jasonisnotapalindrome to NOT be a palindrome")
	}
}
