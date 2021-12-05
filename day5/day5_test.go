package day5_test

import (
	. "ryepup/advent2021/day5"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_test(t *testing.T) {
	n, err := Part1("test.txt")
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
}

func TestParseVent(t *testing.T) {
	vent, err := ParseVent("1,2 -> 3,4")
	assert.Nil(t, err)
	expected := NewVent(1, 2, 3, 4)
	assert.Equal(t, expected, vent)
}

func TestNewVent_vertical(t *testing.T) {
	vent := NewVent(1, 2, 1, 4)
	assert.False(t, vent.IsHorizontal)
	assert.True(t, vent.IsVertical)
}

func TestNewVent_horizontal(t *testing.T) {
	vent := NewVent(1, 2, 3, 2)
	assert.True(t, vent.IsHorizontal)
	assert.False(t, vent.IsVertical)
}

func TestNewVent_diagonal(t *testing.T) {
	vent := NewVent(1, 2, 3, 4)
	assert.False(t, vent.IsHorizontal)
	assert.False(t, vent.IsVertical)
}

func TestPart1_input(t *testing.T) {
	n, err := Part1("input.txt")
	assert.Nil(t, err)
	assert.Equal(t, 6856, n)
}

// func TestPart2_test(t *testing.T) {
// 	n, err := Part2("test.txt")
// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, n)
// }

// func TestPart2_input(t *testing.T) {
// 	n, err := Part2("input.txt")
// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, n)
// }
