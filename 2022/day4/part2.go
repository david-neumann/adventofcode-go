package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartTwo() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	subsetCount := 0

	for sc.Scan() {
		line := sc.Text()
		pairs := strings.Split(line, ",")
		elf1 := parseRange(pairs[0])
		elf2 := parseRange(pairs[1])
		if setsOverlap(elf1, elf2) || setsOverlap(elf2, elf1) {
			subsetCount++
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(subsetCount)
}

func setsOverlap(set1, set2 map[int]struct{}) bool {
	for key := range set1 {
		if _, exists := set2[key]; exists {
			return true
		}
	}
	return false
}
