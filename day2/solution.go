package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getBoxIds() []string {
	var lines []string

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func countOcurrence(ocurrenceToCount int, charOcurrences map[rune]int) int {
	total := 0
	alreadyPassed := false

	for _, ocurrence := range charOcurrences {
		if ocurrenceToCount == ocurrence && !alreadyPassed {
			total++
			alreadyPassed = true
		}
	}

	return total
}

func diff(id1, id2 string) (int, []rune) {
	diff := 0
	var equal []rune

	for i := range id1 {
		if diff > 1 {
			return diff, nil
		}

		if id1[i] != id2[i] {
			diff++
		} else {
			equal = append(equal, rune(id1[i]))

		}
	}

	return diff, equal
}

func part1(boxIDs []string) {
	two := 0
	three := 0

	for _, boxID := range boxIDs {
		charOcurrences := map[rune]int{}

		for _, boxIDChar := range boxID {
			charOcurrences[boxIDChar]++
		}

		two += countOcurrence(2, charOcurrences)
		three += countOcurrence(3, charOcurrences)
	}

	println(two * three)
}

func part2(boxIDs []string) {
	for i, boxID := range boxIDs {
		for _, boxIDToCompare := range boxIDs[i+1:] {
			d, same := diff(boxID, boxIDToCompare)

			if d == 1 {
				fmt.Println(string(same))
			}
		}
	}
}

func main() {
	boxIDs := getBoxIds()
	part1(boxIDs)
	part2(boxIDs)
}
