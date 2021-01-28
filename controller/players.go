package controller

import (
	"errors"
	"log"

	"github.com/lucas-engen/WarTelemetry/model/mapobjects"
)

// GetCurrentPlayer function gets the current player
func GetCurrentPlayer(mo []mapobjects.MapObjects) (*mapobjects.MapObjects, error) {
	var player *mapobjects.MapObjects = &mapobjects.MapObjects{}
	var found bool = false

	// Parses all objects in current map
	for _, object := range mo {
		if IsCurrentPlayer(&object) {
			player = &object
			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("player object could not be found")
	}

	return player, nil
}

// GetAlliesColor function gets the color of allied players
func GetAlliesColor(mo []mapobjects.MapObjects) (color string) {

	// Gets current player
	player, err := GetCurrentPlayer(mo)
	if err != nil {
		log.Printf("[ERROR] Failed to get current player: %v", err)
		return
	}

	return player.Color
}

// IsEnemy function checks if the specified map object is an enemy object
func IsEnemy(object *mapobjects.MapObjects, mo []mapobjects.MapObjects) (bool, error) {

	player, err := GetCurrentPlayer(mo)
	if err != nil {
		return false, err
	}

	if object.Color != player.Color {
		return true, nil
	}

	return false, nil
}
