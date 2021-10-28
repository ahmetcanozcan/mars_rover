package terrain

import (
	"app/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyObject struct {
	c common.Location
}

func (d *dummyObject) Location() common.Location {
	return d.c
}

func newDummyObject(c common.Location) *dummyObject {
	return &dummyObject{
		c: c,
	}
}

func (d *dummyObject) SetTerrain(Terrain)            {}
func (d *dummyObject) SetLocation(l common.Location) {}

var _ Object = (*dummyObject)(nil)

func TestTerrain_GetName(t *testing.T) {
	name := "test"
	terrain := New(name, 100, 100)
	assert.Equal(t, name, terrain.GetName())
}

func TestTerrain_checkBoundaries(t *testing.T) {
	terrain := terrainImpl{
		width:  10,
		height: 10,
	}
	tests := []struct {
		arg      common.Location
		expected bool
	}{
		{common.NewLocation(1, 0), true},
		{common.NewLocation(0, 0), true},
		{common.NewLocation(0, 1), true},
		{common.NewLocation(0, 10), true},
		{common.NewLocation(0, 11), false},
		{common.NewLocation(0, 9), true},
		{common.NewLocation(1, 10), true},
		{common.NewLocation(10, 10), true},
		{common.NewLocation(11, 10), false},
		{common.NewLocation(10, 11), false},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, terrain.checkBoundary(test.arg), "arg: %v", test.arg)
	}
}

func TestTerrain_AddObject(t *testing.T) {
	terrain := New("test", 10, 10)

	t.Run("should use object location for inbound locations", func(t *testing.T) {
		inboundLoc := common.NewLocation(5, 5)
		object := newDummyObject(inboundLoc)
		terrain.AddObject(object)
		assert.True(t, object.c.Equals(inboundLoc))
	})

	t.Run("should use default location for outbound locations", func(t *testing.T) {
		outboundLoc := common.NewLocation(15, 15)
		object := newDummyObject(outboundLoc)
		terrain.AddObject(object)
		assert.True(t, object.c.Equals(outboundLoc))
	})
}

func TestTerrain_IsLocationAvailable(t *testing.T) {
	terrain := New("test", 10, 10)

	tests := []struct {
		arg    common.Location
		wanted bool
	}{
		{common.NewLocation(4, 2), true},
		{common.NewLocation(5, 2), true},
		{common.NewLocation(15, 15), false}, // out of bounds
	}

	objectLocations := []common.Location{
		common.NewLocation(0, 0),
		common.NewLocation(4, 4),
		common.NewLocation(1, 2),
	}

	t.Run("returns true for an empty terrain", func(t *testing.T) {
		for _, test := range tests {
			assert.Equal(t, test.wanted, terrain.IsLocationAvailable(test.arg))
		}
	})

	for _, Location := range objectLocations {
		tests = append(tests, struct {
			arg    common.Location
			wanted bool
		}{Location, false})
	}

	for _, c := range objectLocations {
		terrain.AddObject(newDummyObject(c))
	}

	t.Run("returns false for an occupied terrain", func(t *testing.T) {
		for _, test := range tests {
			assert.Equal(t, test.wanted, terrain.IsLocationAvailable(test.arg))
		}
	})

}
