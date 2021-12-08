package day7_test

import (
	"fmt"
	. "ryepup/advent2021/day7"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	testCases := map[string]int{"test.txt": 37, "input.txt": 345197}
	for path, expected := range testCases {
		t.Run(path, func(t *testing.T) {
			if expected == -1 {
				t.Skip("TODO: provide an expected value")
			}
			n, err := Part1(path)
			require.Nil(t, err)
			require.Equal(t, expected, n)
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := map[string]int{"test.txt": 168, "input.txt": 96361606}
	for path, expected := range testCases {
		t.Run(path, func(t *testing.T) {
			if expected == -1 {
				t.Skip("TODO: provide an expected value")
			}
			n, err := Part2(path)
			require.Nil(t, err)
			require.Equal(t, expected, n)
		})
	}
}

func BenchmarkPart2(b *testing.B) {

	for _, par := range []ParallelStrategy{ForLoop, WaitGroup} {
		for _, cache := range []CacheStrategy{NoCache, Mutex, RWMutex, Naive} {
			b.Run(fmt.Sprintf("P:%v C:%v", par, cache), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, err := Part2Opts("input.txt", cache, par)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}
