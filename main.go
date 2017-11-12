package main

import (
	"fmt"

	"github.com/jasonmichels/euler/smallestmultiple"
)

func main() {
	result := smallestmultiple.Naive()
	fmt.Printf("Smallest Multiple: %v \n", result)
}
