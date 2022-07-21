package geographic

type Coordinates struct {
	X int
	Y int
}

// Add returns the Coordinates from adding the provided coordinates to current
func (c Coordinates) Add(coordinates Coordinates) Coordinates {
	return Coordinates{X: c.X + coordinates.X, Y: c.Y + coordinates.Y}
}

// Scale returns the current Coordinates multiplied by scale
func (c Coordinates) Scale(scale int) Coordinates {
	return Coordinates{
		X: c.X * scale,
		Y: c.Y * scale,
	}
}
