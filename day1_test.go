package adventofcode2021

import (
	"testing"
)

func TestPart1(t *testing.T) {
	n, err := Part1("day1.part1.test.txt")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if n != 7 {
		// TODO: is there an assert.equals(7, n)?
		t.Fatalf("Expected %v, got %v", 7, n)
	}
}
