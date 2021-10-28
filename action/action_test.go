package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		s         string
		want      Action
		wantedErr error
	}{
		{
			s:         "M",
			want:      Move,
			wantedErr: nil,
		},
		{
			s:         "L",
			want:      TurnLeft,
			wantedErr: nil,
		},
		{
			s:         "R",
			want:      TurnRight,
			wantedErr: nil,
		},
		{
			s:         "F",
			want:      Invalid,
			wantedErr: ErrInvalidAction,
		},
	}
	for _, test := range tests {
		got, err := FromString(test.s)
		assert.Equal(t, test.want, got)
		assert.Equal(t, test.wantedErr, err)
	}
}
