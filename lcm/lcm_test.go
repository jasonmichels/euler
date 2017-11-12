package lcm

import (
	"testing"
)

func TestLCM24And54(t *testing.T) {
	if result := LCM(54, 24); result != 216 {
		t.Errorf("We received result of: %v", result)
	}
}

func BenchmarkLCM24And54(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LCM(54, 24)
	}
}
