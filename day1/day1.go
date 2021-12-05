package day1

import (
	"ryepup/advent2021/utils"
)

func countIncreases(depths []int) int {
	lastDepth := depths[0]
	increases := 0

	for _, depth := range depths[1:] {
		if depth > lastDepth {
			increases++
		}
		lastDepth = depth
	}
	return increases
}

func Part1(path string) (int, error) {
	depths, err := utils.ReadInts(path)
	if err != nil {
		return 0, err
	}

	return countIncreases(depths), nil
}

func Part2(path string) (int, error) {
	depths, err := utils.ReadInts(path)
	if err != nil {
		return 0, err
	}

	groupDepths := make([]int, 0)
	group := depths[:3]
	groupDepths = append(groupDepths, utils.SumInts(group))

	// rolling windows for the rest
	for _, depth := range depths[3:] {
		group = append(group[1:], depth)
		groupDepths = append(groupDepths, utils.SumInts(group))
	}

	return countIncreases(groupDepths), nil
}
