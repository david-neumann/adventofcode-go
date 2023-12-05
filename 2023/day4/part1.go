package day4

import (
	"bufio"
	"log"
	"os"
)

func PartOne(filepath string) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	totalPoints := 0

	for sc.Scan() {
		winningNums, yourNums := parseLine(sc.Text())
		yourWinningNums := findWinningNums(winningNums, yourNums)
		points := calculatePoints(yourWinningNums)
		totalPoints += points
	}

	return totalPoints
}

func parseLine(line string) (map[int]struct{}, []int) {
	winningNums := make(map[int]struct{})
	var yourNums []int

	return winningNums, yourNums
}

func findWinningNums(winningNums map[int]struct{}, yourNums []int) []int {
	var yourWinningNums []int

	for _, num := range yourNums {
		if _, exists := winningNums[num]; exists {
			yourWinningNums = append(yourWinningNums, num)
		}
	}

	return yourWinningNums
}

func calculatePoints(yourWinningNums []int) int {
	points := 0

	switch len(yourWinningNums) {
	case 0:
		return 0
	case 1:
		return 1
	default:
		points = len(yourWinningNums) * 2
	}

	return points
}

// 1. Parse each line, store winning numbers into a map, your numbers into []int
// 2. Loop through your numbers, check if present in winning numbers map
// 3. Add winning numbers to a new []int
// 4. Check length of winning numbers
//      - if == 1, return 1 point
//      - if > 1, mult by 2 and return points
// 5. Store points for each line in []int
// 6. Add up total points
