package day4

import (
	"bufio"
	"log"
	"os"
)

func PartTwo(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sumOfCards := 0
	copiesOfEachCard := make(map[int]int)

	cardNum := 1
	for sc.Scan() {
		copiesOfEachCard[cardNum]++
		winningNums, yourNums := parseLine(sc.Text()) // winningNums = map[int]struct{} ; yourNums = []int
		wins := 0
		for _, num := range yourNums {
			if _, exists := winningNums[num]; exists {
				wins++
				copiesOfEachCard[cardNum+wins] += copiesOfEachCard[cardNum]
			}
		}
		sumOfCards += copiesOfEachCard[cardNum]
		cardNum++
	}

	return sumOfCards
}

// 1. Iterate through winningNumsPerCard
// 2. For the number of winning numbers, add 1
