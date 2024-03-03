package main

import (
	"bufio"
	"regexp"
	// "encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	red   int
	blue  int
	green int
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
	fmt.Println("Advent of Code 2023: Day 2")
	fmt.Println("===========================")

	data := readFile("2023/data/day_2.txt")

	// Part 1
	fmt.Println("\tPart 1")
	result_1 := partOne(data)
	fmt.Printf("-> %d\n", result_1)
	// Part 2
	fmt.Println("\tPart 2")
	result_2 := partTwo(data)
	fmt.Printf("-> %d\n", result_2)
}

func partOne(data []string) int {

	pattern := regexp.MustCompile("[0-9]+ [a-z]+")

	c := 0

	bad_game := false

	cubes := map[string]int{"red": 12, "blue": 14, "green": 13}

	for _, row := range data {

		games := strings.Split(row, ":")

		id_game, _ := strconv.Atoi(strings.Split(games[0], " ")[1])

		match := pattern.FindAllString(games[1], -1)

		for _, cube := range match {
			cube_split := strings.Split(cube, " ")

			cube_num, _ := strconv.Atoi(cube_split[0])

			if cubes[cube_split[1]] < cube_num {
				bad_game = true
				break
			}

		}

		if !bad_game {
			c += id_game
		}
		bad_game = false
		cubes = map[string]int{"red": 12, "blue": 14, "green": 13}
	}

	return c

}


func partTwo(data []string) int {
	pattern := regexp.MustCompile("[0-9]+ [a-z]+")

	c := 0

	cubes := map[string]int{"red": 0, "blue": 0, "green": 0}

	for _, row := range data {

		games := strings.Split(row, ":")

		match := pattern.FindAllString(games[1], -1)

		for _, cube := range match {
			cube_split := strings.Split(cube, " ")

			cube_num, _ := strconv.Atoi(cube_split[0])

			if cubes[cube_split[1]] < cube_num {
                cubes[cube_split[1]] = cube_num
			}

		}

		c += cubes["red"] * cubes["blue"] * cubes["green"]
		cubes = map[string]int{"red": 0, "blue": 0, "green": 0}
	}

	return c
}
