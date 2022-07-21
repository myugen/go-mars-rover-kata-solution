package geographic_test

import (
	"testing"

	"github.com/myugen/go-mars-rover-kata/geographic"
	"github.com/stretchr/testify/assert"
)

func TestPlane_PositionBaseOn(t *testing.T) {
	type fields struct {
		width  int
		height int
	}
	type args struct {
		currentPosition geographic.Coordinates
		movement        geographic.Coordinates
	}
	var tests = []struct {
		name   string
		fields fields
		args   args
		want   geographic.Coordinates
	}{
		{
			name:   "when starting from point [4, 3] and moving [1, 1] in a planar which size is [5, 4], the final position should be [0, 0]",
			fields: fields{width: 5, height: 4},
			args: args{
				currentPosition: geographic.Coordinates{X: 4, Y: 3},
				movement:        geographic.Coordinates{X: 1, Y: 1},
			},
			want: geographic.Coordinates{X: 0, Y: 0},
		},
		{
			name:   "when starting from point [0, 0] and moving [-1, -1] in a planar which size is [5, 4], the final position should be [4, 3]",
			fields: fields{width: 5, height: 4},
			args: args{
				currentPosition: geographic.Coordinates{X: 0, Y: 0},
				movement:        geographic.Coordinates{X: -1, Y: -1},
			},
			want: geographic.Coordinates{X: 4, Y: 3},
		},
		{
			name:   "when starting from point [4, 0] and moving [1, -1] in a planar which size is [5, 4], the final position should be [0, 3]",
			fields: fields{width: 5, height: 4},
			args: args{
				currentPosition: geographic.Coordinates{X: 4, Y: 0},
				movement:        geographic.Coordinates{X: 1, Y: -1},
			},
			want: geographic.Coordinates{X: 0, Y: 3},
		},
		{
			name:   "when starting from point [0, 3] and moving [-1, 1] in a planar which size is [5, 4], the final position should be [4, 0]",
			fields: fields{width: 5, height: 4},
			args: args{
				currentPosition: geographic.Coordinates{X: 0, Y: 3},
				movement:        geographic.Coordinates{X: -1, Y: 1},
			},
			want: geographic.Coordinates{X: 4, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plane := geographic.Plane{Width: tt.fields.width, Height: tt.fields.height}
			assert.Equal(t, tt.want, plane.PositionBaseOn(tt.args.currentPosition, tt.args.movement))
		})
	}
}
