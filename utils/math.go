package utils

import "math"

// Vector2D is a bidimensional vector
type Vector2D struct {
	X float64
	Y float64
}

// Hypotenuse function returns the hypotenuse of a triangle
func Hypotenuse(a, b float64) float64 {
	return math.Sqrt((a * a) + (b * b))
}

// CalculateDistance function calculates the distance between two points
func CalculateDistance(p1, p2 Vector2D) float64 {
	dx := math.Pow(p2.X-p1.X, 2)
	dy := math.Pow(p2.Y-p1.Y, 2)

	return math.Sqrt(dx + dy)
}
