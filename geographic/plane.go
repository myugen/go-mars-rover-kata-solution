package geographic

import "github.com/myugen/go-mars-rover-kata/utils"

type Plane struct {
	Width  int
	Height int
}

// PositionBaseOn returns correct Coordinates having in count the spherical plane
func (p Plane) PositionBaseOn(currentPosition, projectedCoordinates Coordinates) Coordinates {
	coordinates := currentPosition.Add(projectedCoordinates)
	return Coordinates{
		X: p.keepValueInRangeFor(coordinates.X, p.Width),
		Y: p.keepValueInRangeFor(coordinates.Y, p.Height),
	}
}

func (Plane) keepValueInRangeFor(value, max int) int {
	return utils.KeepValueInRangeFor(value, 0, max)
}
