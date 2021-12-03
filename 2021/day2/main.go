package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_INPUTS = 1024

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var moves [MAX_INPUTS]string
	for i := 0; scanner.Scan(); i++ {
		moves[i] = scanner.Text()
	}

	fmt.Printf("Product of horizontal/depth positions: %d\n", productPositions(moves[:]))
	fmt.Printf("Product of horizontal/depth/aim positions: %d\n", productAimedPositions(moves[:]))
}

type position struct {
	direction string
	amount    int
}

func productPositions(moves []string) int {
	x := 0
	y := 0

	positions := parseFields(moves)

	for _, position := range positions {
		// exit once you run out of values
		if position.direction == "" {
			break
		}

		switch position.direction {
		case "forward":
			x += position.amount
		case "down":
			y += position.amount
		case "up":
			y -= position.amount
		}
	}

	return x * y
}

func productAimedPositions(moves []string) int {
	x := 0
	y := 0
	aim := 0

	positions := parseFields(moves)

	for _, position := range positions {
		// exit once you run out of values
		if position.direction == "" {
			break
		}

		switch position.direction {
		case "forward":
			x += position.amount
			if aim != 0 {
				y += aim * position.amount
			}
		case "down":
			aim += position.amount
		case "up":
			aim -= position.amount
		}
	}

	return x * y
}

func parseFields(moves []string) []position {
	var positions = make([]position, MAX_INPUTS)

	for i, move := range moves {
		// exit once you run out of values
		if move == "" {
			break
		}

		fields := strings.Split(move, " ")

		direction := fields[0]
		amount, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		positions[i] = position{direction, amount}
	}

	return positions
}
