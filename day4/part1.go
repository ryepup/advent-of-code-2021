package day4

import (
	"fmt"
	"ryepup/advent2021/utils"
	"strconv"
	"strings"
)

/*
You're already almost 1.5km (almost a mile) below the surface of the ocean,
already so deep that you can't see any sunlight. What you can see, however, is a
giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers.
Numbers are chosen at random, and the chosen number is marked on all boards on
which it appears. (Numbers may not appear on all boards.) If all numbers in any
row or any column of a board are marked, that board wins. (Diagonals don't
count.)

The submarine has a bingo subsystem to help passengers (currently, you and the
giant squid) pass the time. It automatically generates a random order in which
to draw numbers and a random set of boards (your puzzle input). For example:

    7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

    22 13 17 11  0
    8  2 23  4 24
    21  9 14 16  7
    6 10  3 18  5
    1 12 20 15 19

    3 15  0  2 22
    9 18 13 17  5
    19  8  7 25 23
    20 11 10 24  4
    14 21 16 12  6

    14 21 17 24  4
    10 16 15  9 19
    18  8 23 26 20
    22 11 13  6  5
    2  0 12  3  7

After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no
winners, but the boards are marked as follows (shown here adjacent to each other
to save space):

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
    8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
    6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
    1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still
no winners:

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
    8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
    6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
    1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

Finally, 24 is drawn:

    22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
    8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
    21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
    6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
    1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

At this point, the third board wins because it has at least one complete row or
column of marked numbers (in this case, the entire top row is marked: 14 21 17
24 4).

The score of the winning board can now be calculated. Start by finding the sum
of all unmarked numbers on that board; in this case, the sum is 188. Then,
multiply that sum by the number that was just called when the board won, 24, to
get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win
first. What will your final score be if you choose that board?
*/
func Part1(path string) (int, error) {
	bingo, err := parseBingo(path)
	if err != nil {
		return 0, err
	}

	for _, n := range bingo.numbers {
		for _, b := range bingo.boards {
			if b.mark(n) {
				fmt.Println(b)
				return n * b.score(), nil
			}
		}
	}

	return 0, fmt.Errorf("no winners")
}

const BINGO_BOARD_SIZE = 5

type bingoBoard struct {
	squares [][]*bingoSquare
}

func (t bingoBoard) String() string {
	var builder strings.Builder
	for _, row := range t.squares {
		for _, square := range row {
			builder.WriteString(fmt.Sprint(square))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

type bingoSquare struct {
	number int
	marked bool
}

func (t bingoSquare) String() string {
	if t.marked {
		return fmt.Sprintf(" _%2v_ ", t.number)
	} else {
		return fmt.Sprintf("  %2v  ", t.number)
	}
}

func (board *bingoBoard) mark(number int) bool {
	for _, row := range board.squares {
		for _, square := range row {
			if square.number == number {
				square.marked = true
			}
		}
	}
	return board.isWinner()
}

func (board *bingoBoard) isWinner() bool {
	return board.hasMarkedRow() || board.hasMarkedColumn()
}

func (board *bingoBoard) hasMarkedRow() bool {
OUTER:
	for _, row := range board.squares {
		for _, square := range row {
			if !square.marked {
				continue OUTER
			}
		}
		return true
	}

	return false
}

func (board *bingoBoard) hasMarkedColumn() bool {
OUTER:
	for col := 0; col < BINGO_BOARD_SIZE; col++ {
		for row := 0; row < BINGO_BOARD_SIZE; row++ {
			square := board.squares[row][col]
			if !square.marked {
				continue OUTER
			}
		}
		return true
	}

	return false
}

func (board *bingoBoard) score() int {
	score := 0
	for _, row := range board.squares {
		for _, square := range row {
			if !square.marked {
				score += square.number
			}
		}
	}
	return score
}

func parseBoard(lines []string) (*bingoBoard, error) {
	if len(lines) != BINGO_BOARD_SIZE {
		return nil, fmt.Errorf("boards want %v lines", BINGO_BOARD_SIZE)
	}

	squares := make([][]*bingoSquare, BINGO_BOARD_SIZE)
	for i, line := range lines {
		nums, err := parseNumbers(strings.Fields(line))
		if err != nil {
			return nil, err
		}
		squares[i] = make([]*bingoSquare, BINGO_BOARD_SIZE)
		for j, n := range nums {
			squares[i][j] = &bingoSquare{number: n, marked: false}
		}
	}

	return &bingoBoard{squares}, nil
}

func parseNumbers(raw []string) ([]int, error) {
	results := make([]int, len(raw))
	for i, s := range raw {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		results[i] = n
	}
	return results, nil
}

func parseBoards(lines []string) ([]*bingoBoard, error) {
	boards := make([]*bingoBoard, 0)
	rawBoard := make([]string, 0)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		rawBoard = append(rawBoard, line)
		if len(rawBoard) == BINGO_BOARD_SIZE {
			board, err := parseBoard(rawBoard)
			if err != nil {
				return nil, err
			}
			boards = append(boards, board)
			rawBoard = make([]string, 0)
		}
	}
	return boards, nil
}

type bingo struct {
	numbers []int
	boards  []*bingoBoard
}

func parseBingo(path string) (*bingo, error) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return nil, err
	}
	numbers, err := parseNumbers(strings.Split(lines[0], ","))
	if err != nil {
		return nil, err
	}
	boards, err := parseBoards(lines[1:])
	if err != nil {
		return nil, err
	}
	return &bingo{numbers, boards}, nil
}
