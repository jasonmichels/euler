package evenfib

// EvenFib Calculate sum of even fibonacci numbers up to before the max
func EvenFib(max int64) int64 {
	// calculate the even fibonacci numbers up to, but not including num
	var sum int64 = 2

	var a int64 = 1
	var b int64 = 2

	for b < max {
		nextFib := a + b

		// if the next fibonacci number is even, add it to the sum
		if nextFib%2 == 0 {
			sum += nextFib
		}

		a = b
		b = nextFib
	}

	return sum
}
