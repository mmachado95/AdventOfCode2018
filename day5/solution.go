package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func getPolymer() string {
	var polymer string

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		polymer = scanner.Text()
	}

	return polymer
}

func canReact(firstChar, secondChar rune) bool {
	if unicode.IsLower(firstChar) && unicode.IsUpper(secondChar) {
		if firstChar == unicode.ToLower(secondChar) {
			return true
		}
	} else if unicode.IsUpper(firstChar) && unicode.IsLower(secondChar) {
		if secondChar == unicode.ToLower(firstChar) {
			return true
		}
	}

	return false
}

func getTypes(polymer string) map[rune]bool {
	types := map[rune]bool{}

	for _, unit := range polymer {
		if _, ok := types[unicode.ToLower(unit)]; !ok {
			types[unicode.ToLower(unit)] = true
		}
	}

	return types
}

func react(polymer string) string {
	for i, unit := range polymer {
		if i == len(polymer)-1 {
			return polymer
		}

		if canReact(unit, rune(polymer[i+1])) {
			newPolymer := polymer[:i] + polymer[i+2:]
			return react(newPolymer)
		}
	}

	return polymer
}

func remove(t rune, polymer string) string {
	lowerCharRemove := strings.Replace(polymer, string(t), "", -1)
	upperCharRemove := strings.Replace(lowerCharRemove, string(unicode.ToUpper(t)), "", -1)
	return upperCharRemove
}

func main() {
	polymer := getPolymer()

	newPolymer := react(polymer)

	fmt.Println(len(newPolymer))

	types := getTypes(polymer)
	minLength := 50000

	for t := range types {
		newPolymer = remove(t, polymer)
		newLength := len(react(newPolymer))

		if newLength < minLength {
			minLength = newLength
		}
	}

	fmt.Println(minLength)
}
