package rover

import (
	"app/action"
	"app/common"
	"app/terrain"
	"errors"
	"fmt"
)

type Rover interface {
	terrain.Object
	Do(action.Action) error
	String() string
}

var (
	ErrInvalidMovement = errors.New("invalid movement")
)

// Verify interface implementation
var _ Rover = (*roverImpl)(nil)

type roverImpl struct {
	location  common.Location
	direction common.Direction
	terrain   terrain.Terrain
}

func New(location common.Location, direction common.Direction) Rover {
	return &roverImpl{
		location:  location,
		direction: direction,
	}
}

func (r *roverImpl) SetTerrain(t terrain.Terrain) {
	r.terrain = t
}

func (r *roverImpl) Location() common.Location {
	return r.location
}

func (r *roverImpl) SetLocation(c common.Location) {
	r.location = c
}

func (r *roverImpl) String() string {
	return fmt.Sprintf("%d %d %s", r.location.X, r.location.Y, r.direction)
}

func (r *roverImpl) Do(a action.Action) error {
	switch a {
	case action.Move:
		return r.move()
	case action.TurnLeft:
		return r.turnLeft()
	case action.TurnRight:
		return r.turnRight()
	}
	return nil
}

func (r *roverImpl) move() error {
	newLocation := r.location.Copy()
	switch r.direction {
	case common.North:
		newLocation.Y++
	case common.South:
		newLocation.Y--
	case common.East:
		newLocation.X++
	case common.West:
		newLocation.X--
	}
	if !r.terrain.IsLocationAvailable(newLocation) {
		return ErrInvalidMovement
	}
	r.location = newLocation
	return nil
}

func (r *roverImpl) turnLeft() error {
	switch r.direction {
	case common.North:
		r.direction = common.West
	case common.South:
		r.direction = common.East
	case common.East:
		r.direction = common.North
	case common.West:
		r.direction = common.South
	}
	return nil
}

func (r *roverImpl) turnRight() error {
	switch r.direction {
	case common.North:
		r.direction = common.East
	case common.South:
		r.direction = common.West
	case common.East:
		r.direction = common.South
	case common.West:
		r.direction = common.North
	}
	return nil
}
