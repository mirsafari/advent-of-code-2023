package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gameCounter := 1
	sum := 0
	sumOfPowers := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if gamePossible(scanner.Text()) {
			sum += gameCounter
		}

		sumOfPowers += gamePossibleFewestCubes(scanner.Text())
		gameCounter += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one:", sum)
	fmt.Println("Part two:", sumOfPowers)
}

// Part one

const maxRedCubes int = 12
const maxGreenCubes int = 13
const maxBlueCubes int = 14

func gamePossible(line string) bool {

	// Remove metadata about games
	line = line[strings.IndexByte(line, ':')+2:]

	// Extract each game into an array
	sets := strings.Split(line, ";")

	regexPattern := `(\d+) (red|green|blue)`
	re := regexp.MustCompile(regexPattern)

	for _, set := range sets {

		cubeCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		matches := re.FindAllStringSubmatch(set, -1)
		for _, match := range matches {
			count := match[1]
			color := match[2]
			cubeCount[color], _ = strconv.Atoi(count)

		}

		if cubeCount["red"] > maxRedCubes {
			return false
		}

		if cubeCount["blue"] > maxBlueCubes {
			return false
		}

		if cubeCount["green"] > maxGreenCubes {
			return false
		}

	}
	return true
}

// Part two
func gamePossibleFewestCubes(line string) int {

	// Remove metadata about games
	line = line[strings.IndexByte(line, ':')+2:]

	// Extract each game into an array
	sets := strings.Split(line, ";")

	regexPattern := `(\d+) (red|green|blue)`
	re := regexp.MustCompile(regexPattern)

	maxRedInGame := 1
	maxGreenInGame := 1
	maxBlueInGame := 1

	for _, set := range sets {

		cubeCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		matches := re.FindAllStringSubmatch(set, -1)
		for _, match := range matches {
			count := match[1]
			color := match[2]
			cubeCount[color], _ = strconv.Atoi(count)

			if cubeCount["red"] > maxRedInGame {
				maxRedInGame = cubeCount["red"]
			}

			if cubeCount["green"] > maxGreenInGame {
				maxGreenInGame = cubeCount["green"]
			}

			if cubeCount["blue"] > maxBlueInGame {
				maxBlueInGame = cubeCount["blue"]
			}
		}
	}
	return maxRedInGame * maxGreenInGame * maxBlueInGame
}
