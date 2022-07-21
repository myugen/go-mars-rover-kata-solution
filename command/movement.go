package command

type Movement int8

const (
	FORWARD  Movement = 1
	BACKWARD Movement = -1
)

var movementsToNames = map[Movement]string{
	FORWARD:  "FORWARD",
	BACKWARD: "BACKWARD",
}

func (m Movement) String() string {
	return movementsToNames[m]
}

func (m Movement) Value() int8 {
	return int8(m)
}

func (m Movement) TurnModifier() int8 {
	return 0
}

func (m Movement) MoveModifier() int8 {
	return m.Value()
}
