package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const inputFile = "2023/day2/input.txt"

func PartOne() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sum := 0

	for id := 1; sc.Scan(); id++ {
		isPossible := isGamePossible(sc.Text())
		if isPossible {
			sum += id
		}
	}

	fmt.Println("Sum: ", sum)
}

func isGamePossible(line string) bool {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	regExp := regexp.MustCompile(`\d+\s+\w+`)
	matches := regExp.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		split := strings.Split(match[0], " ")
		count, _ := strconv.Atoi(split[0])
		color := split[1]
		if count > maxCubes[color] {
			return false
		}
	}

	return true
}

// 1. For each line, can use the index + 1 to get the ID of the game
// 2. Each iteration, need to keep track of each color total
// 3. end of the iteration, compare the totals to the max possible
// 4. if possible, add the ID to the sum
