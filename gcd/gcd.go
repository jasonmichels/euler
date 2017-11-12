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

// Recursive GCD found using recursion. About twice as slow but same result
func Recursive(x, y int) int {
	if y == 0 {
		return x
	}

	return Recursive(y, x%y)
}
