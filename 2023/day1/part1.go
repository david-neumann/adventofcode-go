package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const inputFile = "2023/day1/input.txt"

func PartOne() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	total := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		firstNum, lastNum, _ := findFirstAndLastNumber(sc.Text())
		combinedNum := combineDigits(firstNum, lastNum)
		total += combinedNum
	}

	fmt.Println("Total:", total)
}

func findFirstAndLastNumber(input string) (int, int, error) {
	var firstNum, lastNum int
	var firstNumFound bool

	for _, char := range input {
		if unicode.IsDigit(char) {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return 0, 0, err
			}

			if !firstNumFound {
				firstNum = num
				firstNumFound = true
			}

			lastNum = num
		}
	}

	return firstNum, lastNum, nil
}

func combineDigits(num1, num2 int) int {
	combinedStr := fmt.Sprintf("%d%d", num1, num2)
	combinedNum, _ := strconv.Atoi(combinedStr)
	return combinedNum
}
