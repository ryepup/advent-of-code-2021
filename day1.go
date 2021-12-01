package adventofcode2021

// TODO: are there namespaces so I can get a "day1" scope?

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// TODO: is there a better composable "stream" interface for reading / parsing?
// cribbed from https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
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

func countIncreases(depths <-chan int) int {
	lastDepth := <-depths
	increases := 0

	for depth := range depths {
		if depth > lastDepth {
			increases++
		}
		lastDepth = depth
	}
	return increases
}

// TODO: are there docstrings?
func Part1(path string) (int, error) {
	depths, err := scan(path)
	if err != nil {
		return 0, err
	}

	return countIncreases(depths), nil
}

// TODO: isn't this defined somewhere in the standard lib?!
func sum(a []int) int {
	n := 0
	for _, x := range a {
		n += x
	}
	return n
}

func Part2(path string) (int, error) {
	depths, err := scan(path)
	if err != nil {
		return 0, err
	}

	groupDepths := make(chan int)
	go func() {
		defer close(groupDepths)
		// first group
		group := []int{<-depths, <-depths, <-depths}
		groupDepths <- sum(group)

		// rolling windows for the rest
		for depth := range depths {
			group = append(group[1:], depth)
			groupDepths <- sum(group)
		}
	}()

	return countIncreases(groupDepths), nil
}

/*
Misc learning notes from today:

- you can pass go channels around as send/recv only types chan<- / <-chan
- vscode go plugin is pretty good, lots of "do what I want" features already
  with no extra config
- defer is pretty sweet
- channels can be used like C# IEnumerable or python yield, but you need to add
  concurrency to do so and be careful about when you close the channel
- a little surprised to have zero problems with closures or pointers
- the stdlib seems a little bare bones, hoping I just haven't found what to
  import yet
*/
