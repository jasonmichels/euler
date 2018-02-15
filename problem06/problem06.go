package problem06

// SumSquares Square each number up to n and add them together
func SumSquares(n int) int {
	var result int

	for i := 1; i <= n; i++ {
		result += (i * i)
	}

	return result
}

// SumSquaresChannel Square each number up to n and add them together
func SumSquaresChannel(n int, channel chan int) {
	var result int

	for i := 1; i <= n; i++ {
		result += (i * i)
	}

	channel <- result
}

// SquareSums Sum all the numbers and then square the final number
func SquareSums(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum

}

// SquareSumsChannel Sum all the numbers and then square the final number
func SquareSumsChannel(n int, channel chan int) {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i
	}

	channel <- sum * sum

}

// SumSquareDiff Subtract the sum of the squares, from the square of the sums
func SumSquareDiff() int {
	// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
	return SquareSums(100) - SumSquares(100)
}

// SumSquareDiffChannel Subtract the sum of the squares, from the square of the sums
func SumSquareDiffChannel() int {
	sqSumC := make(chan int)
	sumSqC := make(chan int)

	go SquareSumsChannel(100, sqSumC)
	go SumSquaresChannel(100, sumSqC)

	a, b := <-sqSumC, <-sumSqC
	// Interestingly, using channels in this instance is 10 times slower
	return a - b
}
