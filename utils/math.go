package utils

import "math"

// Hypotenuse function returns the hypotenuse of a triangle
func Hypotenuse(a, b float64) float64 {
	return math.Sqrt((a * a) + (b * b))
}
