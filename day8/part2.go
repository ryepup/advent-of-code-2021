package day8

import "strings"

/*
Through a little deduction, you should now be able to determine the remaining
digits. Consider again the first example above:

	acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb
cdbaf

After some careful analysis, the mapping between signal wires and segments only
make sense in the following configuration:

     dddd
    e    a
    e    a
     ffff
    g    b
    g    b
     cccc

So, the unique signal patterns would correspond to the following digits:

    acedgfb: 8
    cdfbe: 5
    gcdfa: 2
    fbcad: 3
    dab: 7
    cefabd: 9
    cdfgeb: 6
    eafb: 4
    cagedb: 0
    ab: 1

Then, the four digits of the output value can be decoded:

    cdfeb: 5
    fcadb: 3
    cdfeb: 5
    cdbaf: 3

Therefore, the output value for this entry is 5353.

Following this same process for each entry in the second, larger example above,
the output value of each entry can be determined:

    fdgacbe cefdb cefbgd gcbe: 8394
    fcgedb cgb dgebacf gc: 9781
    cg cg fdcagb cbg: 1197
    efabcd cedba gadfec cb: 9361
    gecf egdcabf bgf bfgea: 4873
    gebdcfa ecba ca fadegcb: 8418
    cefg dcbef fcge gbcadfe: 4548
    ed bcgafe cdgba cbgef: 1625
    gbdfcae bgc cg cgb: 8717
    fgae cfgab fg bagce: 4315

Adding all of the output values in this larger example produces 61229.

For each entry, determine all of the wire/segment connections and decode the
four-digit output values. What do you get if you add up all of the output
values?
*/
func Part2(path string) (int, error) {
	entries, err := parseEntries(path)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, entry := range entries {
		mapping, err := entry.findMapping()
		if err != nil {
			return 0, err
		}
		display, err := entry.findDisplay(mapping)
		if err != nil {
			return 0, err
		}
		sum += display
	}

	return sum, nil
}

func (e entry) findMapping() (wireSegmentMapping, error) {
	result := make(map[displaySegment]displaySegment)

	// all things are possible initially
	candidates := NewPotentialSolution()

	// every entry has 10 numbers, one for each valid display
	digitSignals := make(map[int]signalPattern)
	signalDigits := make(map[signalPattern]int)

	for _, size := range []int{2, 3, 4, 7} {
		digit := sizeDigitMap[size][0]
		for _, signal := range e.signals {
			if len(signal) != size {
				continue
			}
			digitSignals[digit] = signal
			signalDigits[signal] = digit
			candidates.require(signal, digit)
		}
	}

	/********************************************************
	 * _EVERYTHING BELOW HERE IS HOT DYSFUNCTIONAL GARBAGE_ *
	 ********************************************************/

	// loop on the rest until we find a solution
	for len(digitSignals) != 10 {
		for _, signal := range e.signals {
			if signalDigits[signal] > 0 {
				continue
			}

			digits := sizeDigitMap[len(signal)]
			fits := make([]int, 0)
			for _, digit := range digits {
				if _, ok := digitSignals[digit]; ok {
					continue
				}
				if candidates.fits(signal, digit) {
					fits = append(fits, digit)
				}
			}
			if len(fits) == 1 {
				digit := fits[0]
				digitSignals[digit] = signal
				signalDigits[signal] = digit
				candidates.require(signal, digit)
			}
		}
	}

	// if signals, ok := signalsByLength[5]; ok {
	// 	for _, d := range []int{2, 3, 5} {
	// 		dst := knownSegmentsByDigit[d]
	// 		for _, signal := range signals {
	// 			candidates.possible(signal, dst...)
	// 		}
	// 	}
	// }
	// if signals, ok := signalsByLength[6]; ok {
	// 	for _, d := range []int{0, 6, 9} {
	// 		dst := knownSegmentsByDigit[d]
	// 		for _, signal := range signals {
	// 			candidates.possible(signal, dst...)
	// 		}
	// 	}
	// }

	//var segmentHistograms [7][7]int
	// for _, signal := range e.signals {
	// 	signalSegments := signal.asDisplaySegments()
	// 	n := len(signal)

	// 	switch n {
	// 	case 2:
	// 		candidates.require(signalSegments, B, C)
	// 	case 3:
	// 		candidates.require(signalSegments, A, B, C)
	// 	case 4:
	// 		candidates.require(signalSegments, F, G, B, C)
	// 	case 7:
	// 		candidates.require(signalSegments, A, B, C, D, E, F, G)

	// 		// case 5:
	// 		// 	// could be a few digits, vote for all them
	// 		// 	for _, d := range []int{2, 3, 5} {
	// 		// 		for _, p := range knownSegmentsByDigit[d] {
	// 		// 			for _, s := range signalSegments {
	// 		// 				segmentHistograms[p][s]++
	// 		// 			}
	// 		// 		}
	// 		// 	}
	// 		// case 6:
	// 		// 	// could be a few digits, vote for all them
	// 		// 	for _, d := range []int{0, 6, 9} {
	// 		// 		for _, p := range knownSegmentsByDigit[d] {
	// 		// 			for _, s := range signalSegments {
	// 		// 				segmentHistograms[p][s]++
	// 		// 			}
	// 		// 		}
	// 		// 	}
	// 	}
	// }

	return result, nil
}

func (e entry) findDisplay(mapping wireSegmentMapping) (int, error) {
	return 0, nil
}

type potentialSolution map[displaySegment][]displaySegment

func NewPotentialSolution() *potentialSolution {
	return &potentialSolution{
		A: {A, B, C, D, E, F, G},
		B: {A, B, C, D, E, F, G},
		C: {A, B, C, D, E, F, G},
		D: {A, B, C, D, E, F, G},
		E: {A, B, C, D, E, F, G},
		F: {A, B, C, D, E, F, G},
		G: {A, B, C, D, E, F, G},
	}
}

func (ps *potentialSolution) require(signal signalPattern, digit int) {
	segments := knownSegmentsByDigit[digit]
	candidates := signal.asDisplaySegments()
	// put them in the right segments
	for _, s := range segments {
		(*ps)[s] = intersectDs((*ps)[s], candidates)
	}
	// drop these candidates from the other destinations, they _must_ be here
	for _, s := range differenceDs(allDs, segments) {
		(*ps)[s] = differenceDs((*ps)[s], candidates)
	}
}

func (ps *potentialSolution) fixed() map[displaySegment]displaySegment {
	result := make(map[displaySegment]displaySegment)
	for k, options := range *ps {
		if len(options) == 1 {
			result[k] = options[0]
		}
	}
	return result
}

func (ps *potentialSolution) fits(signal signalPattern, digit int) bool {
	candidates := signal.asDisplaySegments()
	// which segments would be on for this digit?
	segments := knownSegmentsByDigit[digit]
	fixed := ps.fixed()
	used := make(map[displaySegment]bool)
	for _, s := range segments {
		if ds, ok := fixed[s]; ok && signal.contains(ds) {
			used[ds] = true
		}

		// are those signals in the candidate list?
		common := intersectDs((*ps)[s], candidates)
		if len(common) == 0 {
			return false
		}
		// remove it from the list!
		candidates = differenceDs(candidates, common)

	}
	return true
}

func (s signalPattern) contains(ds displaySegment) bool {
	return strings.ContainsRune(string(s), rune(ds))
}

/*
fdcge

*/

func intersectDs(a, b []displaySegment) []displaySegment {
	seen := make(map[displaySegment]int)

	for _, a := range a {
		seen[a]++
	}
	for _, b := range b {
		seen[b]++
	}

	results := make([]displaySegment, 0)
	for k, v := range seen {
		if v > 1 {
			results = append(results, k)
		}
	}
	return results
}

func differenceDs(a, b []displaySegment) []displaySegment {
	diff := make(map[displaySegment]bool)

	for _, a := range a {
		diff[a] = true
	}
	for _, b := range b {
		delete(diff, b)
	}

	keys := make([]displaySegment, 0, len(diff))
	for ds := range diff {
		keys = append(keys, ds)
	}
	return keys
}

type wireSegmentMapping map[displaySegment]displaySegment

// https://en.wikipedia.org/wiki/Seven-segment_display#/media/File:7_Segment_Display_with_Labeled_Segments.svg\
type displaySegment rune

const (
	A displaySegment = 'a'
	B                = 'b'
	C                = 'c'
	D                = 'd'
	E                = 'e'
	F                = 'f'
	G                = 'g'
)

var allDs = []displaySegment{A, B, C, D, E, F, G}

// possible displayed numbers for each number of input signals
var sizeDigitMap = map[int][]int{
	2: {1},
	3: {7},
	4: {4},
	5: {2, 3, 5},
	6: {0, 6, 9},
	7: {8},
}

var knownSegmentsByDigit = map[int][]displaySegment{
	0: {A, B, C, D, E, F},
	1: {B, C},
	2: {A, B, G, E, D},
	3: {A, B, G, C, D},
	4: {F, G, B, C},
	5: {A, F, G, C, D},
	6: {A, F, G, C, D, E},
	7: {A, B, C},
	8: {A, B, C, D, E, F, G},
	9: {A, F, G, B, C, D},
}

func (s signalPattern) asDisplaySegments() []displaySegment {
	results := make([]displaySegment, len(s))
	for i, s := range s {
		results[i] = asDisplaySegment(s)
	}
	return results
}

func asDisplaySegment(s rune) displaySegment {
	switch s {
	case 'a':
		return A
	case 'b':
		return B
	case 'c':
		return C
	case 'd':
		return D
	case 'e':
		return E
	case 'f':
		return F
	case 'g':
		return G
	}
	panic("shouldn't be possible?")
}
