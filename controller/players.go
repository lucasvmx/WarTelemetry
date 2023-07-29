package controller

import (
	"errors"
	"strings"

	"github.com/lucasvmx/WarTelemetry/logger"
	"github.com/lucasvmx/WarTelemetry/model/mapobjects"
)

var currentPlayerColor string

// GetCurrentPlayer function gets the current player
func GetCurrentPlayer(mo []mapobjects.MapObjects) (*mapobjects.MapObjects, error) {
	var player *mapobjects.MapObjects = &mapobjects.MapObjects{}
	var found bool = false

	// Parses all objects in current map
	for _, object := range mo {
		if IsCurrentPlayer(&object) {
			player = &object
			found = true
			currentPlayerColor = object.Color
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

	if len(currentPlayerColor) > 0 {
		return currentPlayerColor
	}

	// Gets current player
	player, err := GetCurrentPlayer(mo)
	if err != nil {
		logger.LogError("failed to get current player: %v", err)
		return
	}

	currentPlayerColor = player.Color

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

// IsEnemyFighter checks if the specified object is a enemy fighter
func IsEnemyFighter(object *mapobjects.MapObjects, mo []mapobjects.MapObjects) bool {

	if strings.ToLower(object.Icon) != "fighter" {
		return false
	}

	// Get ally color
	allyColor := GetAlliesColor(mo)

	for _, obj := range mo {
		objType := strings.ToLower(obj.Icon)
		if objType == "fighter" && obj.Color != allyColor {
			return true
		}
	}

	return false
}
