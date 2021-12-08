package day7_test

import (
	. "ryepup/advent2021/day7"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testCases := map[string]int{"test.txt": 37, "input.txt": 345197}
	for path, expected := range testCases {
		t.Run(path, func(t *testing.T) {
			if expected == -1 {
				t.Skip("TODO: provide an expected value")
			}
			n, err := Part1(path)
			require.Nil(t, err)
			require.Equal(t, expected, n)
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := map[string]int{"test.txt": 168, "input.txt": 96361630}
	for path, expected := range testCases {
		t.Run(path, func(t *testing.T) {
			if expected == -1 {
				t.Skip("TODO: provide an expected value")
			}
			n, err := Part2(path)
			require.Nil(t, err)
			require.Equal(t, expected, n)
		})
	}
}
