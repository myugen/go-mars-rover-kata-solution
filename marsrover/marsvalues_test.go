package marsrover_test

import (
	"testing"

	"github.com/myugen/go-mars-rover-kata/command"
	"github.com/myugen/go-mars-rover-kata/geographic"
	"github.com/myugen/go-mars-rover-kata/marsrover"
	"github.com/stretchr/testify/assert"
)

func TestMarsRover_Receive_Should(t *testing.T) {
	type fields struct {
		planet        geographic.Plane
		startingPoint geographic.Coordinates
		facingTo      geographic.Direction
	}
	type args struct {
		signals []command.Signal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   marsrover.Status
	}{
		{
			name: "add the provided coordinates to the current position",
			fields: fields{
				planet:        geographic.Plane{Width: 5, Height: 5},
				startingPoint: geographic.Coordinates{X: 2, Y: 4},
				facingTo:      geographic.EAST,
			},
			args: args{
				signals: []command.Signal{
					command.FORWARD,
					command.FORWARD,
					command.LEFT,
					command.BACKWARD,
					command.BACKWARD,
					command.LEFT,
					command.FORWARD,
					command.FORWARD,
					command.FORWARD,
					command.RIGHT,
					command.RIGHT,
					command.RIGHT,
				},
			},
			want: marsrover.Status{
				Position: geographic.Coordinates{X: 1, Y: 1},
				Facing:   geographic.SOUTH,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marsRover := marsrover.DeployInto(tt.fields.planet, tt.fields.startingPoint, tt.fields.facingTo)
			marsRover.Receive(tt.args.signals...)

			assert.Equal(t, tt.want, marsRover.Status())
		})
	}
}
