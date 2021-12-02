package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	var depths [2048]int

	scanner := bufio.NewScanner(file)

	for i:=0; scanner.Scan(); i++ {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		depths[i] = depth
	}

	fmt.Printf("Number of increases: %d\n", countDepthIncreases(depths[:]))
	fmt.Printf("Number of windowed increases: %d\n", countWindowedIncreases(depths[:]))
}

func countDepthIncreases(depths []int) int {
	numIncreases := 0
	prev := depths[0]
	for i:=1; i < len(depths); i++ {
		if prev < depths[i] {
			numIncreases++
		}

		prev = depths[i]
	}

	return numIncreases
}

func countWindowedIncreases(depths []int) int {
	numDepths := len(depths)
	if numDepths < 4 {
		// list must be 4 or greater to build at least two three-part windows
		return 0
	}

	numIncreases := 0
	prev := depths[0] + depths[1] + depths[2]
	for i:=1; i < numDepths - 2; i++ {
		current := depths[i] + depths[i+1] + depths[i+2]
		if prev < current {
			numIncreases++
		}

		prev = current
	}

	return numIncreases
}