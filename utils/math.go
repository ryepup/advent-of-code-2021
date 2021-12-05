package utils

/*
Sum the provides numbers

This is a pre-generics solution

TODO: isn't this defined somewhere in the standard lib?!
*/
func SumInts(a []int) int {
	n := 0
	for _, x := range a {
		n += x
	}
	return n
}

// builtin math.Min is all float64s
func MinInt(numbers ...int) int {
	winner := numbers[0]
	for _, n := range numbers[1:] {
		if n < winner {
			winner = n
		}
	}
	return winner
}

// builtin math.Max is all float64s
func MaxInt(numbers ...int) int {
	winner := numbers[0]
	for _, n := range numbers[1:] {
		if n > winner {
			winner = n
		}
	}
	return winner
}

func AbsInt(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}
