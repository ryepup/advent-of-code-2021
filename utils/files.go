package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

/*
Read a file line-by-line and parse each line to an int

TODO: is there a better composable "stream" interface for reading / parsing?

cribbed from https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
*/
func ReadLines(path string) (<-chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// TODO: does returning a channel make sense, or is it better to pass one in?
	results := make(chan string)
	go func() {
		defer file.Close()
		defer close(results)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			results <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return results, nil
}

/*
Read a file line-by-line and parses each line to an int
*/
func ParseInts(path string) (<-chan int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	results := make(chan int)
	go func() {
		defer close(results)

		for line := range lines {
			// TODO: what's the right thing to do with errors in a channel
			// producer?
			n, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Could not parse %s: %v", line, err)
			}
			results <- n
		}

	}()

	return results, nil
}
