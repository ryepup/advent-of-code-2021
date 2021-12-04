package day4_test

import (
	. "ryepup/advent2021/day4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_test(t *testing.T) {
	n, err := Part1("test.txt")
	assert.Nil(t, err)
	assert.Equal(t, 4512, n)
}

func TestPart1_input(t *testing.T) {
	n, err := Part1("input.txt")
	assert.Nil(t, err)
	assert.Equal(t, 2745, n)
}
