package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectionFromString(t *testing.T) {
	tests := []struct {
		input  string
		output Direction
		err    error
	}{
		{input: "N", output: North, err: nil},
		{input: "n", output: -1, err: ErrInvalidDirection},
		{input: "S", output: South, err: nil},
		{input: "s", output: -1, err: ErrInvalidDirection},
		{input: "E", output: East, err: nil},
		{input: "e", output: -1, err: ErrInvalidDirection},
		{input: "W", output: West, err: nil},
		{input: "w", output: -1, err: ErrInvalidDirection},
		{input: "X", output: -1, err: ErrInvalidDirection},
	}

	for _, test := range tests {
		direction, err := NewDirectionFromString(test.input)
		assert.Equal(t, test.output, direction)
		assert.Equal(t, test.err, err)
	}
}
