package common

import "errors"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var ErrInvalidDirection = errors.New("invalid direction")

var stringDirection = map[string]Direction{
	"N": North,
	"E": East,
	"S": South,
	"W": West,
}

var directionToString = reverseDirectionMap(stringDirection)

func NewDirectionFromString(s string) (Direction, error) {
	if d, ok := stringDirection[s]; ok {
		return d, nil
	}
	return -1, ErrInvalidDirection
}

func (d Direction) String() string {
	if s, ok := directionToString[d]; ok {
		return s
	}
	return ""
}

func reverseDirectionMap(m map[string]Direction) map[Direction]string {
	n := make(map[Direction]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}
