package day7

import (
	"fmt"
	"ryepup/advent2021/utils"
	"sync"
)

/*
The crabs don't seem interested in your proposed solution. Perhaps you
misunderstand crab engineering?

As it turns out, crab submarine engines don't burn fuel at a constant rate.
Instead, each change of 1 step in horizontal position costs 1 more unit of fuel
than the last: the first step costs 1, the second step costs 2, the third step
costs 3, and so on.

As each crab moves, moving further becomes more expensive. This changes the best
horizontal position to align them all on; in the example above, this becomes 5:

    Move from 16 to 5: 66 fuel
    Move from 1 to 5: 10 fuel
    Move from 2 to 5: 6 fuel
    Move from 0 to 5: 15 fuel
    Move from 4 to 5: 1 fuel
    Move from 2 to 5: 6 fuel
    Move from 7 to 5: 3 fuel
    Move from 1 to 5: 10 fuel
    Move from 2 to 5: 6 fuel
    Move from 14 to 5: 45 fuel

This costs a total of 168 fuel. This is the new cheapest possible outcome; the
old alignment position (2) now costs 206 fuel instead.

Determine the horizontal position that the crabs can align to using the least
fuel possible so they can make you an escape route! How much fuel must they
spend to align to that position?
*/
func Part2(path string) (int, error) {
	return Part2Opts(path, NoCache, WaitGroup) // fastest approach on my machine
}

func Part2Opts(path string, cache CacheStrategy, proc ParallelStrategy) (int, error) {
	if cache == Naive && proc != ForLoop {
		return 0, fmt.Errorf("invalid strategy combination")
	}

	positions, err := utils.ReadIntCsv(path)
	if err != nil {
		return 0, err
	}
	maxPosition := utils.MaxInt(positions...)
	solutions := make([]int, maxPosition+1)
	costFunction := makeFuelStrategy(cache)
	processor := makeProcessor(proc)

	processor(positions, maxPosition, costFunction, solutions)

	return utils.MinInt(solutions...), nil
}

func rawFuelCost(distance int) int {
	// https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF
	return (distance * (distance + 1)) / 2
}

type CacheStrategy int

const (
	NoCache CacheStrategy = iota
	Mutex
	RWMutex
	Naive
)

type costFunction = func(int) int

func makeFuelStrategy(strategy CacheStrategy) costFunction {
	if strategy == NoCache {
		return rawFuelCost
	}

	fuelCostCache := make(map[int]int)

	if strategy == Naive {
		return func(distance int) int {
			cached, ok := fuelCostCache[distance]
			if ok {
				return cached
			}
			cost := rawFuelCost(distance)
			fuelCostCache[distance] = cost
			return cost
		}
	}

	if strategy == Mutex {
		var mutex sync.Mutex
		return func(distance int) int {
			mutex.Lock()
			defer mutex.Unlock()
			cached, ok := fuelCostCache[distance]
			if ok {
				return cached
			}
			cost := rawFuelCost(distance)
			fuelCostCache[distance] = cost
			return cost
		}
	}

	if strategy == RWMutex {
		var rwMutex sync.RWMutex
		return func(distance int) int {
			rwMutex.RLock()
			cached, ok := fuelCostCache[distance]
			rwMutex.RUnlock()
			if ok {
				return cached
			}

			rwMutex.Lock()
			defer rwMutex.Unlock()

			cached, ok = fuelCostCache[distance]
			if ok {
				return cached
			}

			cost := rawFuelCost(distance)
			fuelCostCache[distance] = cost
			return cost
		}
	}
	return nil
}

type ParallelStrategy int

const (
	ForLoop ParallelStrategy = iota
	WaitGroup
)

type processor = func([]int, int, costFunction, []int)

func makeProcessor(strategy ParallelStrategy) processor {
	if strategy == ForLoop {
		return forLoopProcessor
	}
	if strategy == WaitGroup {
		return waitGroupProcessor
	}
	return nil
}

func forLoopProcessor(positions []int, maxPosition int, cost costFunction, solutions []int) {
	for target := 0; target <= maxPosition; target++ {
		fuel := 0
		for _, position := range positions {
			distance := utils.AbsInt(position - target)
			fuel += cost(distance)
		}
		solutions[target] = fuel
	}
}

func waitGroupProcessor(positions []int, maxPosition int, cost costFunction, solutions []int) {
	var wg sync.WaitGroup
	for target := 0; target <= maxPosition; target++ {
		wg.Add(1)
		go func(target int) {
			defer wg.Done()
			fuel := 0
			for _, position := range positions {
				distance := utils.AbsInt(position - target)
				fuel += cost(distance)
			}
			solutions[target] = fuel

		}(target)
	}
	wg.Wait()
}
