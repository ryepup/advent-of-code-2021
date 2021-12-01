package adventofcode2021

// TODO: are there namespaces so I can get a "day1" scope?

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// TODO: are there docstrings?
// cribbed from https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func Part1(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lastDepth := 0
	increases := 0
	initial := true

	for scanner.Scan() {
		line := scanner.Text()
		depth, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		if !initial && depth > lastDepth {
			increases++
		}
		initial = false
		lastDepth = depth

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return increases, nil
}
