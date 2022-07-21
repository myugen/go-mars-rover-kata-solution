package marsrover

import (
	"github.com/myugen/go-mars-rover-kata/command"
	"github.com/myugen/go-mars-rover-kata/geographic"
)

type MarsRover struct {
	plane  geographic.Plane
	status Status
}

func (m *MarsRover) Receive(signals ...command.Signal) {
	for _, signal := range signals {
		m.status.Facing = m.status.Facing.TurnTo(signal.TurnModifier())
		coordinatesToAdd := m.status.Facing.MoveTo(signal.MoveModifier())
		m.status.Position = m.plane.PositionBaseOn(m.status.Position, coordinatesToAdd)
	}

}

func (m MarsRover) Status() Status {
	return m.status
}

func DeployInto(planet geographic.Plane, startingPoint geographic.Coordinates, facingTo geographic.Direction) *MarsRover {
	return &MarsRover{plane: planet, status: Status{Position: startingPoint, Facing: facingTo}}
}
