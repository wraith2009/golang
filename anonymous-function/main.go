// compute_example.go
package main

import (
	"fmt"
	"math"
)

// compute takes a function 'fn' that accepts two float64 arguments and returns a float64.
// It calls the function with fixed arguments 3 and 4, and returns the result.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// Define an anonymous function 'hypot' to compute the hypotenuse using the Pythagorean theorem.
	// It takes two float64 values (x and y) and returns sqrt(x^2 + y^2).
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// Call the 'hypot' function directly with arguments 5 and 12.
	// This will compute sqrt(5^2 + 12^2) = sqrt(25 + 144) = sqrt(169) = 13.
	fmt.Println(hypot(5, 12)) // Output: 13

	// Pass the 'hypot' function as an argument to 'compute'.
	// Inside compute, it uses arguments 3 and 4.
	// So it computes sqrt(3^2 + 4^2) = sqrt(9 + 16) = sqrt(25) = 5.
	fmt.Println(compute(hypot)) // Output: 5

	// Pass the standard library function 'math.Pow' to 'compute'.
	// math.Pow(x, y) computes x raised to the power y.
	// So this will compute 3^4 = 81.
	fmt.Println(compute(math.Pow)) // Output: 81
}
