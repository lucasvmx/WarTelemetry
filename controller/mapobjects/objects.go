package controller

import (
	"strings"

	"github.com/lucas-engen/WarTelemetry/model/mapobjects"
)

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
