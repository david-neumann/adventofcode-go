package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo() {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	total := 0

	for sc.Scan() {
		fmt.Println("Initial:", sc.Text())
		line := convertSpelledNumbers(sc.Text())
		fmt.Println("Converted:", line)
		firstNum, lastNum, _ := findFirstAndLastNumber(line)
		fmt.Printf("First: %d, Last: %d\n", firstNum, lastNum)
		combinedNum := combineDigits(firstNum, lastNum)
		fmt.Println("Combined:", combinedNum)
		total += combinedNum
		fmt.Println("Total:", total)
	}

	fmt.Println("Total:", total)
}

func convertSpelledNumbers(line string) string {
	spelledNumbers := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for spelled, digit := range spelledNumbers {
		line = strings.ReplaceAll(line, spelled, digit)
	}

	return line
}
