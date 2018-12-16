package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Rectangle represents the area of fabric
type Rectangle struct {
	Width  int
	Height int
}

// Claim represents the elfs claims
type Claim struct {
	ID    int
	X     int
	Y     int
	Board Rectangle
}

func getClaimValues(line string) []int {
	re := regexp.MustCompile(`^#(\d+)\s*@\s(\d+),(\d+):\s*(\d+)x(\d+)$`)
	stringValues := re.FindAllStringSubmatch(line, -1)

	id, _ := strconv.Atoi(stringValues[0][1])
	x, _ := strconv.Atoi(stringValues[0][2])
	y, _ := strconv.Atoi(stringValues[0][3])
	w, _ := strconv.Atoi(stringValues[0][4])
	h, _ := strconv.Atoi(stringValues[0][5])

	return []int{id, x, y, w, h}
}

func parseClaims() []Claim {
	var claims []Claim

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := getClaimValues(line)

		rectange := Rectangle{
			Width:  values[3],
			Height: values[4],
		}
		claim := Claim{
			ID:    values[0],
			X:     values[1],
			Y:     values[2],
			Board: rectange,
		}

		claims = append(claims, claim)
	}

	return claims
}

func part1(fabric [1000][1000]int, claims []Claim) {
	for _, claim := range claims {
		for x := claim.X; x < claim.X+claim.Board.Width; x++ {
			for y := claim.Y; y < claim.Y+claim.Board.Height; y++ {
				fabric[x][y]++
			}
		}
	}

	overlap := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if fabric[x][y] > 1 {
				overlap++
			}
		}
	}

	fmt.Println(overlap)
}

func part2(fabric [1000][1000]int, claims []Claim) {
	overlap := [1348]bool{}

	for _, claim := range claims {
		for x := claim.X; x < claim.X+claim.Board.Width; x++ {
			for y := claim.Y; y < claim.Y+claim.Board.Height; y++ {
				if fabric[x][y] > 0 {
					overlap[claim.ID] = true
					overlap[fabric[x][y]] = true
				} else {
					fabric[x][y] = claim.ID
				}
			}
		}
	}

	for i := 1; i < 1348; i++ {
		if !overlap[i] {
			fmt.Println(i)
		}
	}

}

func main() {
	fabric := [1000][1000]int{}
	claims := parseClaims()

	part2(fabric, claims)
}
