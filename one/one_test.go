package one

import (
	"testing"
)

func TestNextDivisible(t *testing.T) {
	if NextDivisible(999, 5) != 995 {
		t.Error("We have incorrect next divisble number")
	}
}

func TestProjectOneCalculated(t *testing.T) {
	result := ProblemOneCalculated(999, 3, 5)
	if result != 233168 {
		t.Errorf("Expected: %v, Received: %v \n", 233168, result)
	}
}

// BenchmarkProjectOneCalculated Took just 2000ns
func BenchmarkProjectOneCalculated(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ProblemOneCalculated(999, 3, 5)
	}
}

func TestProjectOneLooped(t *testing.T) {
	if ProblemOneLooped(999, 3, 5) != 233168 {
		t.Error("Unexpected return from project one looped")
	}
}

// BenchmarkProjectOneLooped Took 14000ns and got slow quick
func BenchmarkProjectOneLooped(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ProblemOneLooped(999, 3, 5)
	}
}

// BenchmarkProjectOneCalculatedLargeNum Took 2000ns, not much change
func BenchmarkProjectOneCalculatedLargeNum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ProblemOneCalculated(999999999, 3, 5)
	}
}

// BenchmarkProjectOneLoopedLargeNum Takes about 15 seconds on my laptop
func BenchmarkProjectOneLoopedLargeNum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ProblemOneLooped(999999999, 3, 5)
	}
}

// BenchmarkProjectOneCalculatedLargestNum Took about 2078ns, not much change from smaller numbers
func BenchmarkProjectOneCalculatedLargestNum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ProblemOneCalculated(999999999999, 3, 5)
	}
}

// BenchmarkProjectOneLoopedLargestNum DONT RUN THIS ON A LAPTOP!!! Mine quit before finishing after hundreds of seconds
// func BenchmarkProjectOneLoopedLargestNum(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		ProblemOneLooped(999999999999, 3, 5)
// 	}
// }
