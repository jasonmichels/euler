package main

import (
	"fmt"

	"github.com/jasonmichels/euler/evenfib"
)

func main() {
	sum := evenfib.EvenFib(4000000)
	fmt.Printf("Sum: %v \n", sum)
}
