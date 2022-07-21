package geographic_test

import (
	"testing"

	"github.com/myugen/go-mars-rover-kata/command"
	"github.com/myugen/go-mars-rover-kata/geographic"
	"github.com/stretchr/testify/assert"
)

func TestDirection_TurnTo(t *testing.T) {
	type arg struct {
		turn int8
	}
	testcases := []struct {
		name      string
		direction geographic.Direction
		arg       arg
		want      geographic.Direction
	}{
		{
			name:      "from NORTH should return EAST when turning to RIGHT",
			direction: geographic.NORTH,
			arg:       arg{turn: command.RIGHT.Value()},
			want:      geographic.EAST,
		},
		{
			name:      "from NORTH should return WEST when turning to LEFT",
			direction: geographic.NORTH,
			arg:       arg{turn: command.LEFT.Value()},
			want:      geographic.WEST,
		},
		{
			name:      "from SOUTH should return WEST when turning to RIGHT",
			direction: geographic.SOUTH,
			arg:       arg{turn: command.RIGHT.Value()},
			want:      geographic.WEST,
		},
		{
			name:      "from SOUTH should return EAST when turning to LEFT",
			direction: geographic.SOUTH,
			arg:       arg{turn: command.LEFT.Value()},
			want:      geographic.EAST,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.direction.TurnTo(tt.arg.turn)
			assert.Equal(t, tt.want.String(), got.String())
		})
	}
}

func TestDirection_MoveTo(t *testing.T) {
	type arg struct {
		movements []int8
	}
	testcases := []struct {
		name      string
		direction geographic.Direction
		arg       arg
		want      geographic.Coordinates
	}{
		{
			name:      "forward from NORTH should return the projected coordinates",
			direction: geographic.NORTH,
			arg:       arg{movements: []int8{command.FORWARD.Value(), command.FORWARD.Value()}},
			want:      geographic.Coordinates{X: 0, Y: -2},
		},
		{
			name:      "backward from NORTH should return the projected coordinates",
			direction: geographic.NORTH,
			arg:       arg{movements: []int8{command.BACKWARD.Value(), command.BACKWARD.Value()}},
			want:      geographic.Coordinates{X: 0, Y: 2},
		},
		{
			name:      "forward from EAST should return the projected coordinates",
			direction: geographic.EAST,
			arg:       arg{movements: []int8{command.FORWARD.Value(), command.FORWARD.Value()}},
			want:      geographic.Coordinates{X: 2, Y: 0},
		},
		{
			name:      "backward from EAST should return the projected coordinates",
			direction: geographic.EAST,
			arg:       arg{movements: []int8{command.BACKWARD.Value(), command.BACKWARD.Value()}},
			want:      geographic.Coordinates{X: -2, Y: 0},
		},
		{
			name:      "forward from SOUTH should return the projected coordinates",
			direction: geographic.SOUTH,
			arg:       arg{movements: []int8{command.FORWARD.Value(), command.FORWARD.Value()}},
			want:      geographic.Coordinates{X: 0, Y: 2},
		},
		{
			name:      "backward movement from SOUTH should return the projected coordinates",
			direction: geographic.SOUTH,
			arg:       arg{movements: []int8{command.BACKWARD.Value(), command.BACKWARD.Value()}},
			want:      geographic.Coordinates{X: 0, Y: -2},
		},
		{
			name:      "forward from WEST should return the projected coordinates",
			direction: geographic.WEST,
			arg:       arg{movements: []int8{command.FORWARD.Value(), command.FORWARD.Value()}},
			want:      geographic.Coordinates{X: -2, Y: 0},
		},
		{
			name:      "backward movement from WEST should return the projected coordinates",
			direction: geographic.WEST,
			arg:       arg{movements: []int8{command.BACKWARD.Value(), command.BACKWARD.Value()}},
			want:      geographic.Coordinates{X: 2, Y: 0},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.direction.MoveTo(tt.arg.movements...)
			assert.Equal(t, tt.want, got)
		})
	}
}
