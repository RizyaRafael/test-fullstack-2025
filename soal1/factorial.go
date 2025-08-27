package main

import (
	"fmt"
	"math"
)

func FactorialMult(n int) int {
	if n == 0 {
		return 1
	}
	factorial := 1
	for i := n; i > 0; i-- {
		factorial *= i
	}
	squared := math.Pow(2, float64(n))
	output := math.Ceil(float64(factorial) / squared)
	return int(output)
}

func main() {
	result := FactorialMult(5)
	fmt.Println(result)
}
