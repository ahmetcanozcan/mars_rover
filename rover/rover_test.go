package rover

import (
	"app/action"
	"app/common"
	"app/terrain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockTerrain struct {
	availableFunc func(common.Location) bool
}

var _ terrain.Terrain = (*mockTerrain)(nil)

func (m *mockTerrain) IsLocationAvailable(Location common.Location) bool {
	return m.availableFunc(Location)
}
func (m *mockTerrain) GetName() string {
	return ""
}
func (m *mockTerrain) AddObject(terrain.Object) {
}

func newMockTerrain(availableFunc func(common.Location) bool) *mockTerrain {
	return &mockTerrain{availableFunc: availableFunc}
}

func TestRover_Location(t *testing.T) {
	l := common.Location{X: 1, Y: 1}
	d := common.North
	r := New(l, d)
	assert.True(t, r.Location().Equals(l))
}

func TestRover_SetTerrain(t *testing.T) {
	flag := false
	flagTerrain := newMockTerrain(func(c common.Location) bool {
		flag = true
		return false
	})
	l := common.Location{X: 1, Y: 1}
	d := common.North
	r := New(l, d)
	r.SetTerrain(flagTerrain)
	r.Do(action.Move)
	assert.True(t, flag)
}

func TestRover_move(t *testing.T) {

	t.Run("should move everywhere on all available terrain", func(t *testing.T) {
		acceptTerrain := newMockTerrain(func(Location common.Location) bool {
			return true
		})

		rover := roverImpl{
			terrain:   acceptTerrain,
			location:  common.Location{X: 0, Y: 0},
			direction: common.North,
		}
		tests := []struct {
			direction common.Direction
			expected  common.Location
		}{
			{common.North, common.Location{X: 0, Y: 1}},
			{common.East, common.Location{X: 1, Y: 1}},
			{common.South, common.Location{X: 1, Y: 0}},
			{common.West, common.Location{X: 0, Y: 0}},
		}

		for _, test := range tests {
			rover.direction = test.direction
			err := rover.move()
			assert.True(t, test.expected.Equals(rover.location), "expected %v, got %v", test.expected, rover.location)
			assert.Nil(t, err)
		}
	})

	t.Run("should not move if terrain is not available", func(t *testing.T) {
		rejectTerrain := newMockTerrain(func(Location common.Location) bool {
			return false
		})

		rover := roverImpl{
			terrain:   rejectTerrain,
			location:  common.Location{X: 0, Y: 0},
			direction: common.North,
		}
		tests := []struct {
			direction common.Direction
			expected  common.Location
		}{
			{common.North, common.Location{X: 0, Y: 0}},
			{common.East, common.Location{X: 0, Y: 0}},
			{common.South, common.Location{X: 0, Y: 0}},
			{common.West, common.Location{X: 0, Y: 0}},
		}

		for _, test := range tests {
			rover.direction = test.direction
			err := rover.move()
			assert.True(t, test.expected.Equals(rover.location), "expected %v, got %v", test.expected, rover.location)
			assert.Equal(t, err, ErrInvalidMovement)
		}
	})
}

func TestRover_turn(t *testing.T) {
	rover := roverImpl{
		direction: common.North,
	}
	tests := []struct {
		expected common.Direction
		action   func() error
	}{
		{
			expected: common.West,
			action:   rover.turnLeft,
		},
		{
			expected: common.South,
			action:   rover.turnLeft,
		},
		{
			expected: common.East,
			action:   rover.turnLeft,
		},
		{
			expected: common.North,
			action:   rover.turnLeft,
		},
		{
			expected: common.East,
			action:   rover.turnRight,
		},
		{
			expected: common.South,
			action:   rover.turnRight,
		},
	}

	for _, test := range tests {
		err := test.action()
		assert.Nil(t, err)
		assert.Equal(t, test.expected, rover.direction)
	}
}
