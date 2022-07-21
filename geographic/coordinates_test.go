package geographic_test

import (
	"testing"

	"github.com/myugen/go-mars-rover-kata/geographic"
	"github.com/stretchr/testify/assert"
)

func TestCoordinates_Add_Should(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		coordinates geographic.Coordinates
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   geographic.Coordinates
	}{
		{
			name:   "return coordinates resulting from adding the provided coordinate to current coordinate",
			fields: fields{X: 2, Y: 4},
			args:   args{coordinates: geographic.Coordinates{X: 4, Y: 1}},
			want:   geographic.Coordinates{X: 6, Y: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coordinates := geographic.Coordinates{X: tt.fields.X, Y: tt.fields.Y}
			assert.Equal(t, tt.want, coordinates.Add(tt.args.coordinates))
		})
	}
}
