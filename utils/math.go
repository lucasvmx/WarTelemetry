package utils

import "math"

// Hypotenuse function returns the hypotenuse of a triangle
func Hypotenuse(a, b float64) float64 {
	return math.Sqrt((a * a) + (b * b))
}

func IsTriangleRect(base, height float64) bool {

	tan := height / base
	angle := math.Atan(tan)

	return angle == math.Pi/2
}
