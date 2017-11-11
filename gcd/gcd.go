package gcd

// GCD greatest common divisor
func GCD(x int, y int) int {

	for y != 0 {
		temp := y
		y = x % y
		x = temp
	}

	return x
}
