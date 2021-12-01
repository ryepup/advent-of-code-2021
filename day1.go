package adventofcode2021

// TODO: are there namespaces so I can get a "day1" scope?

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// TODO: is there a better composable "stream" interface for reading / parsing?
func scan(path string) (<-chan int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// TODO: does returning a channel make sense, or is it better to pass one in?
	results := make(chan int)
	go func() {
		defer file.Close()
		defer close(results)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			// TODO: what's the right thing to do with errors in a channel producer?
			n, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Could not parse %s: %v", line, err)
			}
			results <- n
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return results, nil
}

// TODO: are there docstrings?
// cribbed from https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func Part1(path string) (int, error) {
	depths, err := scan(path)
	if err != nil {
		return 0, err
	}

	lastDepth := <-depths
	increases := 0

	for depth := range depths {
		if depth > lastDepth {
			increases++
		}
		lastDepth = depth
	}

	return increases, nil
}
