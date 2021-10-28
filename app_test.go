package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockWriter struct {
	content string
}

func newMockWriter() *mockWriter {
	return &mockWriter{
		content: "",
	}
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.content += string(p)
	return len(p), nil
}

var _ io.Writer = (*mockWriter)(nil)

func TestStartApp(t *testing.T) {
	tests := []struct {
		input  string
		wanted string
	}{
		{
			input: buildText(
				`5 5
				1 2 N
				LMLMLMLMM
				3 3 E
				MMRMMRMRRM
			`),
			wanted: buildText(
				`1 3 N
				5 1 E
			`),
		},
		{
			input: buildText(
				`6 6
				2 2 N
				RMMMR
				2 2 N
				RMMMMMRMMMMMM
			`),
			wanted: buildText(
				`5 2 S
				4 0 S
			`),
		},
	}

	for _, test := range tests {
		wr := newMockWriter()
		rd := strings.NewReader(test.input)
		startApp(rd, wr)
		assert.Equal(t, test.wanted, wr.content)
	}
}

func buildText(txt string) string {
	return strings.Replace(txt, "\t", "", -1)
}
