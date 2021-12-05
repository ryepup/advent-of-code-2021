package utils

import "strconv"

/*
via strconv.Atoi
*/
func ParseNumbers(raw []string) ([]int, error) {
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
