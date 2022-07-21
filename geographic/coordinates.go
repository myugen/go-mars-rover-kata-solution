package geographic

type Coordinates struct {
	X int
	Y int
}

func (c Coordinates) Add(coordinates Coordinates) Coordinates {
	return Coordinates{X: c.X + coordinates.X, Y: c.Y + coordinates.Y}
}

func (c Coordinates) Scale(scale int) Coordinates {
	return Coordinates{
		X: c.X * scale,
		Y: c.Y * scale,
	}
}
