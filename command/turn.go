package command

type Turn int8

const (
	RIGHT Turn = 1
	LEFT  Turn = -1
)

var turnsToNames = map[Turn]string{
	LEFT:  "LEFT",
	RIGHT: "RIGHT",
}

func (t Turn) String() string {
	return turnsToNames[t]
}

func (t Turn) TurnModifier() int8 {
	return int8(t)
}

func (t Turn) MoveModifier() int8 {
	return 0
}
