package action

import "errors"

type Action int

const (
	Invalid Action = iota
	Move
	TurnLeft
	TurnRight
)

var (
	ErrInvalidAction = errors.New("invalid action")
)

var stringToAction = map[string]Action{
	"M": Move,
	"L": TurnLeft,
	"R": TurnRight,
}

func FromString(s string) (Action, error) {
	a, ok := stringToAction[s]
	if !ok {
		return Invalid, ErrInvalidAction
	}
	return a, nil
}
