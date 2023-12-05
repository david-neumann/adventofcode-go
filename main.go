package main

import (
	"fmt"

	"github.com/david-neumann/adventofcode-go/2023/day4"
)

const (
	inputFile = "2023/day4/input.txt"
	testFile  = "2023/day4/test_input.txt"
)

func main() {
	answer := day4.PartOne(inputFile)
	fmt.Println("Answer:", answer)
	day4.PartOne(inputFile)
}
