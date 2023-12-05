package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "2022/day4/input.txt"

func PartOne() {
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
		if isSubset(elf1, elf2) || isSubset(elf2, elf1) {
			subsetCount++
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(subsetCount)
}

func parseRange(input string) map[int]struct{} {
	set := make(map[int]struct{})
	parts := strings.Split(input, "-")
	if len(parts) == 2 {
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			set[i] = struct{}{}
		}
	}

	return set
}

func isSubset(set1, set2 map[int]struct{}) bool {
	for key := range set1 {
		if _, exists := set2[key]; !exists {
			return false
		}
	}
	return true
}
