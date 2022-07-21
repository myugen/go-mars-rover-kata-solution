package geographic

import (
	"github.com/myugen/go-mars-rover-kata/utils"
)

type Direction int8

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

var (
	directionsToNames = map[Direction]string{
		NORTH: "NORTH",
		EAST:  "EAST",
		SOUTH: "SOUTH",
		WEST:  "WEST",
	}
	directionsToProjectedCoordinates = map[Direction]Coordinates{
		NORTH: {
			X: 0,
			Y: -1,
		},
		EAST: {
			X: 1,
			Y: 0,
		},
		SOUTH: {
			X: 0,
			Y: 1,
		},
		WEST: {
			X: -1,
			Y: 0,
		},
	}
)

func (d Direction) String() string {
	return directionsToNames[d]
}

func (d Direction) Value() int8 {
	return int8(d)
}

// TurnTo returns the corresponding Direction due to turns
func (d Direction) TurnTo(turns ...int8) Direction {
	var totalTurns int8
	for _, turn := range turns {
		totalTurns += turn
	}

	currentDirectionValue := d.Value()
	nextDirectionValue := d.keepValuesInRangeFor(currentDirectionValue + totalTurns)
	return Direction(nextDirectionValue)
}

// MoveTo returns the projected Coordinates scaled by movements
func (d Direction) MoveTo(movements ...int8) Coordinates {
	var totalMovements int8
	for _, movement := range movements {
		totalMovements += movement
	}

	if totalMovements < int8(0) {
		currentDirectionValue := d.Value()
		backwardTurnModifier := int8(2)
		backwardDirectionValue := d.keepValuesInRangeFor(currentDirectionValue + backwardTurnModifier)
		projectedCoordinates := directionsToProjectedCoordinates[Direction(backwardDirectionValue)]
		return projectedCoordinates.Scale(int(-totalMovements))
	}
	return directionsToProjectedCoordinates[d].Scale(int(totalMovements))
}

func (d Direction) keepValuesInRangeFor(value int8) int8 {
	return int8(utils.KeepValueInRangeFor(int(value), 0, 4))
}
