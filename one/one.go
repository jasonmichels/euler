package one

import (
	"github.com/jasonmichels/euler/lcm"
)

// Project Euler Problem 1
// Find the sum of all the multiples of 3 or 5 below 1000

// NextDivisible Find the next divisible number from the largest and the multiple
func NextDivisible(largest int, divisibleBy int) int {
	var divisible int
	for i := largest; i > 0; i-- {
		if i%divisibleBy == 0 {
			divisible = i
			break
		}
	}

	return divisible
}

func sumOfDivisible(largest int, multiple int, channel chan int) {
	nextDivisible := NextDivisible(largest, multiple)
	divisibleBy := nextDivisible / multiple
	channel <- divisibleBy * (nextDivisible + multiple) / 2
}

// ProblemOneCalculated Solve project euler problem 1 building up a calculation. Fastest way to solve problem
func ProblemOneCalculated(largest int, multipleA int, multipleB int) int {

	sumsC := make(chan int)
	go sumOfDivisible(largest, multipleA, sumsC)
	go sumOfDivisible(largest, multipleB, sumsC)

	minusC := make(chan int)
	least := lcm.LCM(multipleA, multipleB)
	go sumOfDivisible(largest, least, minusC)

	a, b, c := <-sumsC, <-sumsC, <-minusC

	return a + b - c
}

// ProblemOneLooped Solve problem one by creating a loop. More concise code, but much slower
func ProblemOneLooped(largest int, multipleA int, multipleB int) int {
	var sum int

	// Could potentially run this in goroutines to speed it up
	for i := 1; i <= largest; i++ {
		if i%multipleA == 0 || i%multipleB == 0 {
			sum += i
		}
	}

	return sum
}