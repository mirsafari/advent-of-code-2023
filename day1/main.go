package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += SumOfSingleLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func SumOfSingleLine(line string) int {

	var numbers []string

	for _, letter := range line {
		_, err := strconv.Atoi(string(letter))

		if err != nil {
			continue
		}

		numbers = append(numbers, string(letter))
	}

	newNumber, _ := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])

	return newNumber
}
