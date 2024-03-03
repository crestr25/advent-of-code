package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Val struct {
    index    int
    oldValue string
    newValue string
}

func readFile(route string) [][]string {
	file, err := os.Open(route)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := csv.NewReader(file).ReadAll()

	if err != nil {
		panic(err)
	}

	return data
}

func main() {
	// New day
	fmt.Println("===========================")
	fmt.Println("Advent of Code 2023: Day 1")
	fmt.Println("===========================")

	data := readFile("2023/data/day_1.txt")

	// Part 1
	fmt.Println("\tPart 1")
	result_1 := partOne(data)
	fmt.Printf("-> %d\n", result_1)
	// Part 2
	fmt.Println("\tPart 2")
	result_2 := partTwo(data)
	fmt.Printf("-> %d\n", result_2)
}

func partOne(rows [][]string) int {
	var first string
	var last string
	var result_list []int

	for _, row := range rows {
		row := string(row[0])
		for i := 0; i < len(row); i++ {
			if _, err := strconv.Atoi(string(row[i])); err == nil {
				if first == "" {
					first = string(row[i])
				}
				last = string(row[i])
			}

		}
		concat_string, _ := strconv.Atoi(first + last)
		result_list = append(result_list, concat_string)

		first = ""
		last = ""
	}

	var result int
	for _, row := range result_list {
		result += row
	}

	return result
}

func partTwo(rows [][]string) int {
	var first string
	var last string
	var result_list []int

	for _, row := range rows {
		vals := numReplacer(row[0])
        
        first = vals[0].newValue
        last = vals[1].newValue

		concat_string, _ := strconv.Atoi(first + last)
		result_list = append(result_list, concat_string)

		fmt.Println(concat_string)
		first = ""
		last = ""
	}

	var result int
	for _, row := range result_list {
		result += row
	}

	return result
}

func numReplacer(s string) []Val {
	intMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	lowestValStr := Val{index: 100}
	highestValStr := Val{index: 0}
	lowestValDig := Val{index: 0}
	highestValDig := Val{index: 0}

    var vals []Val

	for k, v := range intMap {
		if strings.Contains(s, k) {
			if ind := strings.Index(s, k); ind != -1 {
				if lowestValStr.index > ind {
					lowestValStr = Val{index: ind, oldValue: k, newValue: v}
				}
			}
			if ind := strings.LastIndex(s, k); ind != -1 {
				if highestValStr.index < ind {
					highestValStr = Val{index: ind, oldValue: k, newValue: v}
				}
			}

		}
	}

    for i := 0; i < len(s); i++ {
        if _, err := strconv.Atoi(string(s[i])); err == nil {
            if lowestValDig.newValue == "" {
                lowestValDig.newValue = string(s[i])
                lowestValDig.index = i
            }
            highestValDig.newValue = string(s[i])
            highestValDig.index = i
        }

    }
    if lowestValStr.index < lowestValDig.index {
        vals = append(vals, lowestValStr)
    } else {
        vals = append(vals, lowestValDig)
    }
    
    if highestValStr.index > highestValDig.index {
        vals = append(vals, highestValStr)
    } else {
        vals = append(vals, highestValDig)
    }

	return vals
}
