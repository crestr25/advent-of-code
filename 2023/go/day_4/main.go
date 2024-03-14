package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	id             int
	numbers        []int
	winningNumbers []int
}

func readFile(route string) []string {
	file, err := os.Open(route)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}

func main() {
	// New day
	fmt.Println("===========================")
	fmt.Println("Advent of Code 2023: Day 4")
	fmt.Println("===========================")

	data := readFile("2023/data/day_4.txt")

	// Part 1
	fmt.Println("\tPart 1")
	result_1 := partOne(data)
	fmt.Printf("-> %d\n", result_1)
	// Part 2
	fmt.Println("\tPart 2")
	result_2 := partTwo(data)
	fmt.Printf("-> %d\n", result_2)
}

func splitRow(row string) Row {
	//
	rowSplit := strings.Split(row, ": ")

	// Id side
	id := strings.Trim(strings.Split(strings.Split(rowSplit[0], ": ")[0], " ")[1], " ")
	idInt, _ := strconv.Atoi(id)
	// idInt = idInt - 1

	// Number side
	numStrings := strings.Split(rowSplit[1], " | ")
	currentNumbers := strings.Split(strings.Trim(string(numStrings[0]), " "), " ")
	winningNumbers := strings.Split(strings.Trim(string(numStrings[1]), " "), " ")

	var intNumbers []int
	var intWinning []int

	for _, number := range currentNumbers {
		if number == "" {
			continue
		}
		n, _ := strconv.Atoi(strings.Trim(string(number), " "))
		intNumbers = append(intNumbers, n)
	}

	for _, number := range winningNumbers {
		if number == "" {
			continue
		}
		n, _ := strconv.Atoi(strings.Trim(string(number), " "))
		intWinning = append(intWinning, n)
	}

	return Row{id: idInt, numbers: intNumbers, winningNumbers: intWinning}
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func partOne(data []string) int {

	var results []int

	for _, line := range data {

		row := splitRow(line)

		c := -1

		for _, number := range row.numbers {
			if s := intInSlice(number, row.winningNumbers); s {
				c += 1
			}
		}

		if c > -1 {
			result := int(math.Pow(2, float64(c)))
			results = append(results, result)
		}
	}

	var i int
	for _, result := range results {
		i += result
	}

	return i
}

func partTwo(data []string) int {

	// var results []int
	result := 0
	total := make(map[int]int)
	total[1] = 0

	c := 0

	for i, line := range data {

		c = i + 1
		row := splitRow(line)
		times := 1
		times = times + total[c]

		add(total, i+1)
		for j := 0; j < times; j++ {
			c = i + 2
			for _, number := range row.numbers {
				if s := intInSlice(number, row.winningNumbers); s {
					add(total, c)
					c += 1
				}
			}

		}
	}
    
    for _, v := range total{
        result += v 
    }

	return result 
}


func add(dict map[int]int, c int) {
	if _, ok := dict[c]; ok {
		dict[c] = dict[c] + 1
	} else {
		dict[c] = 1
	}
}
