package day3_test

import (
	. "ryepup/advent2021/day3"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	n, err := Part1("test.txt")
	assert.Nil(t, err)
	assert.Equal(t, 198, n)
}

func TestPart1_Input(t *testing.T) {
	n, err := Part1("input.txt")
	assert.Nil(t, err)
	assert.Equal(t, 3309596, n)
}

func TestPart2(t *testing.T) {
	n, err := Part2("test.txt")
	assert.Nil(t, err)
	assert.Equal(t, 230, n)
}
