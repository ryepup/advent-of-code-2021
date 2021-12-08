package day7

import (
	"math"
	"ryepup/advent2021/utils"
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
	positions, err := utils.ReadIntCsv(path)
	if err != nil {
		return 0, err
	}

	target := mean(positions)
	fuel := 0
	for _, position := range positions {
		distance := utils.AbsInt(position - target)
		fuel += fuelCost(distance)
	}

	return fuel, nil
}

func mean(nums []int) int {
	sum := utils.SumInts(nums)
	n := len(nums)
	return int(math.Round(float64(sum) / float64(n)))
}

func fuelCost(distance int) int {
	/*
		1 -> 1: 1
		2 -> 3: 1 + 2
		3 -> 6: 1 + 2 + 3
		4 -> 10: 1 + 2 + 3 + 4
	*/
	// TODO: there's probably a cheaper / clever math trick to do thi
	if distance == 1 || distance == 0 {
		return distance
	} else {
		return distance + fuelCost(distance-1)
	}
}
