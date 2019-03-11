package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z*z-x) < 1e-6 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
