package utils

import (
	"bufio"
	"os"
	"strconv"
)

/*
Read a file line-by-line and parse each line to an int

TODO: is there a better composable "stream" interface for reading / parsing?

cribbed from https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
*/
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

/*
Read a file line-by-line and parses each line to an int
*/
func ParseInts(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	results := make([]int, len(lines))
	for i, line := range lines {
		results[i], err = strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
