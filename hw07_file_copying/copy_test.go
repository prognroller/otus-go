package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name             string
		fromPath, toPath string
		limit, offset    int64
	}{
		{name: "Negative offset", fromPath: "testdata/input.txt", toPath: "out.txt", offset: -1, limit: 0},
		{name: "Negative limit", fromPath: "testdata/input.txt", toPath: "out.txt", offset: 0, limit: -1},
		{name: "Offset larger than file", fromPath: "testdata/input.txt", toPath: "out.txt", offset: 10000, limit: 0},
		{name: "Unsupported file", fromPath: "testdata/input1.txt", toPath: "out.txt", offset: 0, limit: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Copy(test.fromPath, test.toPath, test.offset, test.limit)
			require.Error(t, err)
		})
	}
}
