package lcm

import (
	"github.com/jasonmichels/euler/gcd"
)

// LCM Find the Least Common Multiple
func LCM(a int, b int, integers ...int) int {
	result := a * b / gcd.GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
