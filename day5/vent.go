package day5

import (
	"fmt"
	"regexp"
	"ryepup/advent2021/utils"
)

type vent struct {
	start, end                           point
	IsHorizontal, IsVertical, IsDiagonal bool
}

func (v vent) String() string {
	return fmt.Sprintf("%v -> %v", v.start, v.end)
}

func NewVent(x1, y1, x2, y2 int) vent {
	return vent{
		start:        NewPoint(x1, y1),
		end:          NewPoint(x2, y2),
		IsHorizontal: y1 == y2,
		IsVertical:   x1 == x2,
		IsDiagonal:   utils.AbsInt(x2-x1) == utils.AbsInt(y2-y1),
	}
}

/*
returns the points in the path of this vent
*/
func (v vent) Path() []point {
	results := []point{v.start}

	minX := utils.MinInt(v.start.x, v.end.x)
	maxX := utils.MaxInt(v.start.x, v.end.x)
	minY := utils.MinInt(v.start.y, v.end.y)
	maxY := utils.MaxInt(v.start.y, v.end.y)

	if v.IsHorizontal {
		y := v.start.y
		for x := minX + 1; x < maxX; x++ {
			results = append(results, point{x, y})
		}
	} else if v.IsVertical {
		x := v.start.x
		for y := minY + 1; y < maxY; y++ {
			results = append(results, point{x, y})
		}
	} else if v.IsDiagonal {
		length := utils.AbsInt(v.end.x - v.start.x)
		xf := 1
		yf := 1
		if v.end.x < v.start.x {
			xf = -1
		}
		if v.end.y < v.start.y {
			yf = -1
		}
		for i := 1; i < length; i++ {
			results = append(results, point{v.start.x + (i * xf), v.start.y + (i * yf)})
		}
	}
	return append(results, v.end)
}

type point struct{ x, y int }

func (p point) String() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func NewPoint(x, y int) point {
	return point{x, y}
}

/*
Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where
x1,y1 are the coordinates of one end the line segment and x2,y2 are the
coordinates of the other end. These line segments include the points at both
ends. In other words:

    An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
    An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
*/
func parseVents(path string) ([]vent, error) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return nil, err
	}
	results := make([]vent, len(lines))
	for i, line := range lines {
		results[i], err = ParseVent(line)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

var ventRegex = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func ParseVent(line string) (vent, error) {
	captures := ventRegex.FindStringSubmatch(line)
	if captures == nil {
		return vent{}, fmt.Errorf("could not parse Vent from %v", line)
	}
	values, err := utils.ParseNumbers(captures[1:])
	if err != nil {
		return vent{}, err
	}
	return NewVent(values[0], values[1], values[2], values[3]), nil
}

func filterVent(vents []vent, predicate func(vent) bool) []vent {
	results := make([]vent, 0)
	for _, vent := range vents {
		if predicate(vent) {
			results = append(results, vent)
		}
	}
	return results
}
