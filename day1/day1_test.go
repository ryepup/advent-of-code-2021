package day1

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

func TestPart2(t *testing.T) {
	n, err := Part2("day1.part1.test.txt")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if n != 5 {
		t.Fatalf("Expected %v, got %v", 5, n)
	}
}

// figuring out how slices/append works
func TestSlicesAndAppend(t *testing.T) {
	s := make([]int, 0, 3)
	s = append(s, 1, 2, 3)

	if !(s[0] == 1 && s[1] == 2 && s[2] == 3) {
		t.Fatalf("Wrong contents %v", s)
	}
	s = append(s[1:], 4)
	if !(s[0] == 2 && s[1] == 3 && s[2] == 4) {
		t.Fatalf("Wrong contents %v", s)
	}
}

func TestPart2RealInput(t *testing.T) {
	n, err := Part2("day1.part1.txt")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if n != 1262 {
		t.Fatalf("Expected %v, got %v", 1262, n)
	}
}
