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
