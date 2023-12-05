package main

import (
	"fmt"

	"github.com/david-neumann/adventofcode-go/2023/day3"
)

const (
	inputFile = "2023/day3/input.txt"
	testFile  = "2023/day3/test_input.txt"
)

func main() {
	answer, _ := day3.PartTwo(inputFile)
	fmt.Println("Answer:", answer)
}
