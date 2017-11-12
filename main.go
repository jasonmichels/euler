package main

import (
	"fmt"

	"github.com/jasonmichels/euler/palindrome"
)

func main() {
	result := palindrome.LargestPalindrome()
	fmt.Printf("Largest: %v \n", result)
}
