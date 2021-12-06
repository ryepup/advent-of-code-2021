package day6_test

import (
	. "ryepup/advent2021/day6"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	// TODO: fill in expected values
	testCases := map[string]int{"test.txt": -1, "input.txt": -1}
	for path, expected := range testCases {
		t.Run(path, func(t *testing.T) {
			n, err := Part1(path)
			require.Nil(t, err)
			require.Equal(t, expected, n)
		})
	}
}

func TestPart2(t *testing.T) {
	// TODO: fill in expected values
	testCases := map[string]int{"test.txt": -1, "input.txt": -1}
	for path, expected := range testCases {
		t.Run(path, func(t *testing.T) {
			n, err := Part1(path)
			require.Nil(t, err)
			require.Equal(t, expected, n)
		})
	}
}
