package common

type Location struct {
	X int
	Y int
}

func (c Location) Copy() Location {
	return Location{c.X, c.Y}
}

func (c Location) Equals(other Location) bool {
	return c.X == other.X && c.Y == other.Y
}

func NewLocation(x int, y int) Location {
	return Location{x, y}
}
