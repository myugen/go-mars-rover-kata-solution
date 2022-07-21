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
			got := coordinates.Add(tt.args.coordinates)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCoordinates_Scale(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	type args struct {
		scale int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   geographic.Coordinates
	}{
		{
			name:   "return coordinates resulting from multiplying the provided scale to current coordinate",
			fields: fields{X: 2, Y: 4},
			args:   args{scale: 4},
			want:   geographic.Coordinates{X: 8, Y: 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := geographic.Coordinates{X: tt.fields.X, Y: tt.fields.Y}
			got := c.Scale(tt.args.scale)
			assert.Equal(t, tt.want, got)
		})
	}
}
