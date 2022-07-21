package marsrover

import "github.com/myugen/go-mars-rover-kata/geographic"

type Status struct {
	Position geographic.Coordinates
	Facing   geographic.Direction
}
