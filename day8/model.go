package day8

import (
	"fmt"
	"strings"
)

type SignalPattern string

func parseSignalPatterns(s string) ([]SignalPattern, error) {
	signals := strings.Fields(s)
	if len(signals) != 10 {
		return nil, fmt.Errorf("incorrect number of signals in %v", s)
	}
	results := make([]SignalPattern, len(signals))
	for i, s := range signals {
		results[i] = SignalPattern(s)
	}
	return results, nil
}

func (s SignalPattern) Single() segment {
	return segment(s[0])
}

func (s SignalPattern) WithoutPattern(other ...SignalPattern) SignalPattern {
	var b strings.Builder
OUTER:
	for _, r := range s {
		for _, o := range other {
			if strings.ContainsRune(string(o), r) {
				continue OUTER
			}
		}
		// must not have been in any of our "others"
		b.WriteRune(r)

	}
	return SignalPattern(b.String())
}

type SegmentPattern string

func parseSegmentPatterns(s string) ([]SegmentPattern, error) {
	segments := strings.Fields(s)
	if len(segments) != 4 {
		return nil, fmt.Errorf("incorrect number of segments in %v", s)
	}
	results := make([]SegmentPattern, len(segments))
	for i, s := range segments {
		results[i] = SegmentPattern(s)
	}
	return results, nil
}

type segment rune

/*
 aaaa
b    c
b    c
 dddd
e    f
e    f
 gggg
*/
const (
	A segment = 'a'
	B segment = 'b'
	C segment = 'c'
	D segment = 'd'
	E segment = 'e'
	F segment = 'f'
	G segment = 'g'
)

var displayNumbers = map[SegmentPattern]int{
	SegmentPattern("abcefg"):  0,
	SegmentPattern("cf"):      1,
	SegmentPattern("acdeg"):   2,
	SegmentPattern("acdfg"):   3,
	SegmentPattern("bcdf"):    4,
	SegmentPattern("abdfg"):   5,
	SegmentPattern("abdefg"):  6,
	SegmentPattern("acf"):     7,
	SegmentPattern("abcdefg"): 8,
	SegmentPattern("abcdfg"):  9,
}

type wireSegmentMapping map[segment]segment

func (mapping wireSegmentMapping) Display(pattern SegmentPattern) int {
	enabled := make(map[segment]bool)
	for _, r := range pattern {
		enabled[mapping[segment(r)]] = true
	}
	// build a sorted string so it can match w/ display numbers
	var b strings.Builder
	for _, s := range []segment{A, B, C, D, E, F, G} {
		if enabled[s] {
			b.WriteRune(rune(s))
		}
	}
	mappedPattern := SegmentPattern(b.String())
	return displayNumbers[mappedPattern]
}

type entry struct {
	signals         []SignalPattern
	signalsByLength map[int][]SignalPattern
	segments        []SegmentPattern
}

func NewEntry(line string) (entry, error) {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return entry{}, fmt.Errorf("expected only one `|`, got %v", line)
	}
	signals, err := parseSignalPatterns(parts[0])
	if err != nil {
		return entry{}, err
	}
	segments, err := parseSegmentPatterns(parts[1])
	if err != nil {
		return entry{}, err
	}

	signalsByLength := make(map[int][]SignalPattern)
	for _, signal := range signals {
		length := len(signal)
		if v, ok := signalsByLength[length]; ok {
			signalsByLength[length] = append(v, signal)
		} else {
			signalsByLength[length] = []SignalPattern{signal}
		}
	}

	return entry{signals, signalsByLength, segments}, nil
}

func (e entry) BuildMapping() wireSegmentMapping {
	result := make(wireSegmentMapping)

	// these signals must match these segments
	acf := e.signalsByLength[3][0]
	cf := e.signalsByLength[2][0]
	bcdf := e.signalsByLength[4][0]
	abcdefg := e.signalsByLength[7][0]

	// derive more segments via subtraction
	a := acf.WithoutPattern(cf)
	result[a.Single()] = A

	bd := bcdf.WithoutPattern(cf)
	eg := abcdefg.WithoutPattern(bcdf, a)

	// find signal for "3", will only work for one
	acdfg := e.firstSignal(func(s SignalPattern) bool {
		return len(s) == 5 && len(s.WithoutPattern(cf)) == 3
	})
	g := acdfg.WithoutPattern(acf, bd)
	result[g.Single()] = G
	seg_e := eg.WithoutPattern(g)
	result[seg_e.Single()] = E

	// find 2
	acdeg := e.firstSignal(func(s SignalPattern) bool {
		return len(s) == 5 && len(s.WithoutPattern(bcdf)) == 3
	})
	d := acdeg.WithoutPattern(eg, acf)
	result[d.Single()] = D
	b := bd.WithoutPattern(d)
	result[b.Single()] = B
	f := cf.WithoutPattern(acdeg)
	result[f.Single()] = F

	// find the last one
	c := abcdefg.WithoutPattern(a, bd, eg, f)
	result[c.Single()] = C

	return result
}

func (e entry) firstSignal(pred func(s SignalPattern) bool) SignalPattern {
	for _, signal := range e.signals {
		if pred(signal) {
			return signal
		}
	}
	panic("no matching signals found!")
}

func (e entry) Display() int {
	mapping := e.BuildMapping()
	a := mapping.Display(e.segments[0])
	b := mapping.Display(e.segments[1])
	c := mapping.Display(e.segments[2])
	d := mapping.Display(e.segments[3])
	return (a * 1000) + (b * 100) + (c * 10) + d
}
