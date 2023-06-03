package controller

import (
	"strings"

	"github.com/lucasvmx/WarTelemetry/model/mapobjects"
	"github.com/lucasvmx/WarTelemetry/utils"
)

// distanceFactor constains the multiplication factor to be used on distance calculation
const distanceFactor float64 = 66.32072

// IsAircraft function determines if an object is of type aircraft
func IsAircraft(object *mapobjects.MapObjects) bool {

	// Check type
	if object.Type == "aircraft" {
		return true
	}

	return false
}

// IsCurrentPlayer function determines if current object is the current player
func IsCurrentPlayer(object *mapobjects.MapObjects) bool {

	iconName := strings.ToLower(object.Icon)

	// Check object
	if iconName == "player" {
		return true
	}

	return false
}

// CalculateDistanceInKm gets the distance in kilometers between two objects
func CalculateDistanceInKm(obj1, obj2 *mapobjects.MapObjects) float64 {

	// Convert object positions into a 2D vector
	obj1Vector := utils.Vector2D{
		X: float64(obj1.X),
		Y: float64(obj1.Y),
	}

	obj2Vector := utils.Vector2D{
		X: float64(obj2.X),
		Y: float64(obj2.Y),
	}

	// Calculates the distance
	distance := utils.CalculateDistance(obj1Vector, obj2Vector)

	// Convert distance to kilometers
	return distance * distanceFactor
}
