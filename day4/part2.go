package day4

import (
	"fmt"
)

/*
On the other hand, it might be wise to try a different strategy: let the giant
squid win.

You aren't sure how many bingo boards a giant squid could play at once, so
rather than waste time counting its arms, the safe thing to do is to figure out
which board will win last and choose that one. That way, no matter which boards
it picks, it will win for sure.

In the above example, the second board is the last to win, which happens after
13 is eventually called and its middle column is completely marked. If you were
to keep playing until this point, the second board would have a sum of unmarked
numbers equal to 148 for a final score of 148 * 13 = 1924.

Figure out which board will win last. Once it wins, what would its final score
be?
*/
func Part2(path string) (int, error) {
	bingo, err := parseBingo(path)
	if err != nil {
		return 0, err
	}

	winners := make(map[*bingoBoard]bool)

	for _, n := range bingo.numbers {
		for _, b := range bingo.boards {
			if winners[b] {
				continue
			}
			if b.mark(n) {
				winners[b] = true
				if len(winners) == len(bingo.boards) {
					fmt.Println(b)
					return n * b.score(), nil
				}
			}
		}
	}

	return 0, fmt.Errorf("no winners")
}
