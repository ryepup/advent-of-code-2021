package day8_test

import (
	"fmt"
	. "ryepup/advent2021/day8"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testCases := map[string]int{"test.txt": 26, "input.txt": 365}
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
	testCases := map[string]int{"test.txt": 61229, "input.txt": 975706}
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

func TestNewEntry(t *testing.T) {
	entry, err := NewEntry("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	require.Nil(t, err)
	require.NotEmpty(t, entry)
}

func TestSignalPattern_WithoutPattern(t *testing.T) {
	pattern := SignalPattern("fbcad")
	testCases := map[string]string{
		"bc": "fad",
		"c":  "fbad",
		"fd": "bca",
	}

	for other, expected := range testCases {
		t.Run(other, func(t *testing.T) {
			result := pattern.WithoutPattern(SignalPattern(other))
			require.Equal(t, SignalPattern(expected), result)
		})
	}
}

func TestSignalPattern_WithoutPattern_multiple(t *testing.T) {
	pattern := SignalPattern("fbcad")
	result := pattern.WithoutPattern(
		SignalPattern("bc"), SignalPattern("c"), SignalPattern("fd"))

	require.Equal(t, SignalPattern("a"), result)
}

func TestNewEntry_BuildMapping(t *testing.T) {
	entry, err := NewEntry("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	require.Nil(t, err)
	mapping := entry.BuildMapping()
	require.Equal(t, A, mapping[D])
	require.Equal(t, B, mapping[E])
	require.Equal(t, C, mapping[A])
	require.Equal(t, D, mapping[F])
	require.Equal(t, E, mapping[G])
	require.Equal(t, F, mapping[B])
	require.Equal(t, G, mapping[C])
}

func TestNewEntry_Display(t *testing.T) {
	testCases := []struct {
		raw      string
		expected int
	}{
		{
			"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
			5353,
		},
		{
			"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
			8394,
		},
	}
	for i, test := range testCases {
		name := fmt.Sprintf("case %v", i)
		t.Run(name, func(t *testing.T) {
			entry, err := NewEntry(test.raw)
			require.Nil(t, err)
			require.Equal(t, test.expected, entry.Display())
		})
	}
}
