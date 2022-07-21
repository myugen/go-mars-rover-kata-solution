package geographic_test

import (
	"testing"

	"github.com/myugen/go-mars-rover-kata/command"
	"github.com/myugen/go-mars-rover-kata/geographic"
	"github.com/stretchr/testify/assert"
)

func TestDirection_TurnTo(t *testing.T) {
	type arg struct {
		turns []int8
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
			arg:       arg{turns: []int8{command.RIGHT.TurnModifier()}},
			want:      geographic.EAST,
		},
		{
			name:      "from NORTH should return WEST when turning to LEFT",
			direction: geographic.NORTH,
			arg:       arg{turns: []int8{command.LEFT.TurnModifier()}},
			want:      geographic.WEST,
		},
		{
			name:      "from SOUTH should return WEST when turning to RIGHT",
			direction: geographic.SOUTH,
			arg:       arg{turns: []int8{command.RIGHT.TurnModifier()}},
			want:      geographic.WEST,
		},
		{
			name:      "from SOUTH should return EAST when turning to LEFT",
			direction: geographic.SOUTH,
			arg:       arg{turns: []int8{command.LEFT.TurnModifier()}},
			want:      geographic.EAST,
		},
		{
			name:      "from SOUTH should return SOUTH when NO turns",
			direction: geographic.SOUTH,
			arg:       arg{turns: []int8{}},
			want:      geographic.SOUTH,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.direction.TurnTo(tt.arg.turns...)
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
			name:      "forward from NORTH should return the corresponding projected coordinates",
			direction: geographic.NORTH,
			arg:       arg{movements: []int8{command.FORWARD.MoveModifier(), command.FORWARD.MoveModifier()}},
			want:      geographic.Coordinates{X: 0, Y: -2},
		},
		{
			name:      "backward from NORTH should return the corresponding projected coordinates",
			direction: geographic.NORTH,
			arg:       arg{movements: []int8{command.BACKWARD.MoveModifier(), command.BACKWARD.MoveModifier()}},
			want:      geographic.Coordinates{X: 0, Y: 2},
		},
		{
			name:      "forward from EAST should return the corresponding projected coordinates",
			direction: geographic.EAST,
			arg:       arg{movements: []int8{command.FORWARD.MoveModifier(), command.FORWARD.MoveModifier()}},
			want:      geographic.Coordinates{X: 2, Y: 0},
		},
		{
			name:      "backward from EAST should return the corresponding projected coordinates",
			direction: geographic.EAST,
			arg:       arg{movements: []int8{command.BACKWARD.MoveModifier(), command.BACKWARD.MoveModifier()}},
			want:      geographic.Coordinates{X: -2, Y: 0},
		},
		{
			name:      "forward from SOUTH should return the corresponding projected coordinates",
			direction: geographic.SOUTH,
			arg:       arg{movements: []int8{command.FORWARD.MoveModifier(), command.FORWARD.MoveModifier()}},
			want:      geographic.Coordinates{X: 0, Y: 2},
		},
		{
			name:      "backward from SOUTH should return the corresponding projected coordinates",
			direction: geographic.SOUTH,
			arg:       arg{movements: []int8{command.BACKWARD.MoveModifier(), command.BACKWARD.MoveModifier()}},
			want:      geographic.Coordinates{X: 0, Y: -2},
		},
		{
			name:      "forward from WEST should return the corresponding projected coordinates",
			direction: geographic.WEST,
			arg:       arg{movements: []int8{command.FORWARD.MoveModifier(), command.FORWARD.MoveModifier()}},
			want:      geographic.Coordinates{X: -2, Y: 0},
		},
		{
			name:      "backward from WEST should return the corresponding projected coordinates",
			direction: geographic.WEST,
			arg:       arg{movements: []int8{command.BACKWARD.MoveModifier(), command.BACKWARD.MoveModifier()}},
			want:      geographic.Coordinates{X: 2, Y: 0},
		},
		{
			name:      "with NO movement from WEST should return ZERO projected coordinates",
			direction: geographic.WEST,
			arg:       arg{movements: []int8{}},
			want:      geographic.Coordinates{X: 0, Y: 0},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.direction.MoveTo(tt.arg.movements...)
			assert.Equal(t, tt.want, got)
		})
	}
}
