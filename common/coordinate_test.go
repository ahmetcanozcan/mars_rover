package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_Equals(t *testing.T) {
	tests := []struct {
		c    Location
		arg  Location
		want bool
	}{
		{
			c:    Location{X: 1, Y: 2},
			arg:  Location{X: 1, Y: 2},
			want: true,
		},
		{
			c:    Location{X: 1, Y: 2},
			arg:  Location{X: 1, Y: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.c.Equals(tt.arg))
	}
}

func TestLocation_Copy(t *testing.T) {
	t.Run("copies the Location", func(t *testing.T) {
		tests := []struct {
			c    Location
			want Location
		}{
			{
				c:    Location{X: 1, Y: 2},
				want: Location{X: 1, Y: 2},
			},
			{
				c:    Location{X: 1, Y: 2},
				want: Location{X: 1, Y: 2},
			},
		}
		for _, tt := range tests {
			assert.Equal(t, tt.want, tt.c.Copy())
		}
	})

	t.Run("does not modify the original", func(t *testing.T) {
		original := Location{X: 1, Y: 2}
		copy := original.Copy()

		// Modify the copy
		copy.X = 3
		copy.Y = 4

		assert.True(t, original.Equals(Location{X: 1, Y: 2}))
		assert.True(t, copy.Equals(Location{X: 3, Y: 4}))
	})
}
