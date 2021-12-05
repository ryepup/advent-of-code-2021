package day5_test

import (
	"fmt"
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
	assert.False(t, vent.IsDiagonal)
}

func TestNewVent_horizontal(t *testing.T) {
	vent := NewVent(1, 2, 3, 2)
	assert.True(t, vent.IsHorizontal)
	assert.False(t, vent.IsVertical)
	assert.False(t, vent.IsDiagonal)
}

func TestNewVent_askew(t *testing.T) {
	vent := NewVent(1, 2, 3, 5)
	assert.False(t, vent.IsHorizontal)
	assert.False(t, vent.IsVertical)
	assert.False(t, vent.IsDiagonal)
}

func TestNewVent_diagonal(t *testing.T) {
	vent := NewVent(1, 1, 3, 3)
	assert.False(t, vent.IsHorizontal)
	assert.False(t, vent.IsVertical)
	assert.True(t, vent.IsDiagonal)

	vent = NewVent(9, 7, 7, 9)
	assert.False(t, vent.IsHorizontal)
	assert.False(t, vent.IsVertical)
	assert.True(t, vent.IsDiagonal)
}

func Example_vent_Path_diagonal_downright() {
	vent := NewVent(2, 2, 4, 0)
	fmt.Println(vent.Path())
	// Output: [2,2 3,1 4,0]
}

func Example_vent_Path_diagonal_downleft() {
	vent := NewVent(2, 2, 0, 0)
	fmt.Println(vent.Path())
	// Output: [2,2 1,1 0,0]
}

func Example_vent_Path_diagonal_upright() {
	vent := NewVent(0, 0, 2, 2)
	fmt.Println(vent.Path())
	// Output: [0,0 1,1 2,2]
}

func Example_vent_Path_diagonal_upleft() {
	vent := NewVent(8, 0, 0, 8)
	fmt.Println(vent.Path())
	// Output: [8,0 7,1 6,2 5,3 4,4 3,5 2,6 1,7 0,8]
}

func TestPart1_input(t *testing.T) {
	n, err := Part1("input.txt")
	assert.Nil(t, err)
	assert.Equal(t, 6856, n)
}

func TestPart2_test(t *testing.T) {
	n, err := Part2("test.txt")
	assert.Nil(t, err)
	assert.Equal(t, 12, n)
}

func TestPart2_input(t *testing.T) {
	n, err := Part2("input.txt")
	assert.Nil(t, err)
	assert.Equal(t, 20666, n)
}
