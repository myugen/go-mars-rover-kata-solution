package command

type Signal interface {
	TurnModifier() int8
	MoveModifier() int8
}
