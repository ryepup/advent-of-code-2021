package day2

import (
	"fmt"
	"log"
	"ryepup/advent2021/utils"
	"strconv"
	"strings"
)

type operation int

const (
	None operation = iota
	Forward
	Down
	Up
)

type command struct {
	operation operation
	arg       int
}

func ParseCommands(path string) (<-chan command, error) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return nil, err
	}

	result := make(chan command)
	go func() {
		defer close(result)
		for line := range lines {
			cmd, err := parseCommand(line)
			if err != nil {
				log.Fatalf("Could not parse %s: %v", line, err)
			}
			result <- cmd
		}
	}()

	return result, nil
}

func parseCommand(line string) (command, error) {
	bits := strings.Split(line, " ")
	if len(bits) != 2 {
		return command{}, fmt.Errorf("expected two elements in the line, found %v", bits)
	}
	op, err := parseOperation(bits[0])
	if err != nil {
		return command{}, err
	}
	arg, err := strconv.Atoi(bits[1])
	if err != nil {
		return command{}, err
	}

	return command{op, arg}, nil
}

func parseOperation(s string) (operation, error) {
	switch s {
	case "forward":
		return Forward, nil
	case "down":
		return Down, nil
	case "up":
		return Up, nil
	}
	return None, fmt.Errorf("could not parse %s", s)
}
