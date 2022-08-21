package marsrover

import (
	"github.com/myugen/go-mars-rover-kata-solution/command"
	"github.com/myugen/go-mars-rover-kata-solution/geographic"
)

type MarsRover struct {
	plane  geographic.Plane
	status Status
}

// Receive update the Status for every sent Signal
func (m *MarsRover) Receive(signals ...command.Signal) {
	for _, signal := range signals {
		m.status.Facing = m.status.Facing.TurnTo(signal.TurnModifier())
		coordinatesToAdd := m.status.Facing.MoveTo(signal.MoveModifier())
		m.status.Position = m.plane.PositionBaseOn(m.status.Position, coordinatesToAdd)
	}

}

// Status returns the current status
func (m MarsRover) Status() Status {
	return m.status
}

// DeployInto instances a new Mars Rover in a starting point (geographic.Coordinates) of a new planet (geographic.Plane) facing a direction (geographic.Direction)
func DeployInto(planet geographic.Plane, startingPoint geographic.Coordinates, facingTo geographic.Direction) *MarsRover {
	return &MarsRover{plane: planet, status: Status{Position: startingPoint, Facing: facingTo}}
}
