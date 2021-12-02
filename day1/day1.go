package day1

import (
	"ryepup/advent2021/utils"
)

func countIncreases(depths <-chan int) int {
	lastDepth := <-depths
	increases := 0

	for depth := range depths {
		if depth > lastDepth {
			increases++
		}
		lastDepth = depth
	}
	return increases
}

func Part1(path string) (int, error) {
	depths, err := utils.ParseInts(path)
	if err != nil {
		return 0, err
	}

	return countIncreases(depths), nil
}

func Part2(path string) (int, error) {
	depths, err := utils.ParseInts(path)
	if err != nil {
		return 0, err
	}

	groupDepths := make(chan int)
	go func() {
		defer close(groupDepths)
		// first group
		group := []int{<-depths, <-depths, <-depths}
		groupDepths <- utils.SumInts(group)

		// rolling windows for the rest
		for depth := range depths {
			group = append(group[1:], depth)
			groupDepths <- utils.SumInts(group)
		}
	}()

	return countIncreases(groupDepths), nil
}
