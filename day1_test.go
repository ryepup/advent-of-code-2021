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

// TODO: figure out how to use `go run .` to do this, getting a bunch
// of errors like "adventofcode2021 is not a main package" or
// found packages adventofcode2021 (day1.go) and main (main.go) in /workspaces/advent-of-code-2021
func TestPart1RealInput(t *testing.T) {
	n, err := Part1("day1.part1.txt")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if n != 1292 {
		t.Fatalf("Expected %v, got %v", 1292, n)
	}
}
