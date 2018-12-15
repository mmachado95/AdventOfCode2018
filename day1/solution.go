package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func exer1() int {
	freq := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		freq += x
	}

	return freq
}

func main() {
	log.Printf("%d", exer1())
}
