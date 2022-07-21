package geographic

import "github.com/myugen/go-mars-rover-kata/utils"

type Plane struct {
	Width  int
	Height int
}

func (d Plane) PositionBaseOn(currentPosition, movement Coordinates) Coordinates {
	coordinates := currentPosition.Add(movement)
	return Coordinates{
		X: d.keepValueInRangeFor(coordinates.X, d.Width),
		Y: d.keepValueInRangeFor(coordinates.Y, d.Height),
	}
}

func (d Plane) keepValueInRangeFor(value, max int) int {
	return utils.KeepValueInRangeFor(value, 0, max)
}
