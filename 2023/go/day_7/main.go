package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

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
	fmt.Println("Advent of Code 2023: Day 7")
	fmt.Println("===========================")

	data := readFile("2023/data/day_7.txt")

	// Part 1
	fmt.Println("\tPart 1")
	result_1 := partOne(data)
	fmt.Printf("-> %d\n", result_1)
	// Part 2
	fmt.Println("\tPart 2")
	result_2 := partTwo(data)
	fmt.Printf("-> %d\n", result_2)
}

func checkString(hand string) map[rune]int {
	fmt.Println(hand)
	bets := make(map[rune]int)
	for _, data := range hand {
		bets[data] = bets[data] + 1
	}
	return bets
}


func checkStringJoker(hand string) map[rune]int {
	fmt.Println(hand)
	bets := make(map[rune]int)
    var maxKey string
    var maxNumber int
	for _, data := range hand {
        if string(data) != "J"{
            bets[data] = bets[data] + 1
            if bets[data] > maxNumber {
                maxKey = string(data)
                maxNumber = bets[data]
            }
        }
	}
    if strings.Contains(hand, "J") == false {
        return bets
    }

    if len(bets) == 0 {
        bets['J'] = 5
        return bets
    }

    hand = strings.Replace(hand, "J", maxKey, -1)

    bets = checkStringJoker(hand)

	return bets
}

func checkBest(list map[string]int, ini string, sec string) bool {
	for i := 0; i < len(ini); i++ {
		if list[string(ini[i])] < list[string(sec[i])] {
			return true
		}
	}
	return false
}

func partOne(data []string) int {

	ind := map[string]int{"A": 1, "K": 2, "Q": 3, "J": 4, "T": 5, "9": 6, "8": 7, "7": 8, "6": 9, "5": 10, "4": 11, "3": 12, "2": 13}

	results := make(map[int][]string)
	result := 0
	c := 1

	for _, data := range data {
		dataSplit := strings.Split(data, " ")
		bets := checkString(dataSplit[0])

		betLen := len(bets)
		v := make([]int, 0, betLen)

		for _, value := range bets {
			v = append(v, value)
		}
		if betLen == 1 {
			results[1] = slices.Insert(results[1], len(results[1]), data)
		} else if betLen == 2 {
			containNum := slices.Contains(v, 1)
			if containNum == false {
				results[3] = slices.Insert(results[3], len(results[3]), data)
			} else {
				results[2] = slices.Insert(results[2], len(results[2]), data)
			}
		} else if betLen == 3 {
			containNum2 := slices.Contains(v, 2)
			containNum3 := slices.Contains(v, 3)
			if containNum3 == true {
				results[4] = slices.Insert(results[4], len(results[4]), data)
			} else if containNum2 == true {
				results[5] = slices.Insert(results[5], len(results[5]), data)
			} else {
				results[6] = slices.Insert(results[6], len(results[6]), data)
			}

		} else if betLen == 4 {
			results[7] = slices.Insert(results[7], len(results[7]), data)
		} else {
			results[8] = slices.Insert(results[8], len(results[8]), data)
		}

	}
	keys := make([]int, 0, len(results))

	for k := range results {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, i := range keys {
		handList := make([]string, 0, len(results[i]))
		betList := make(map[string]int)
		for _, value := range results[i] {
			dataSplit := strings.Split(value, " ")
			betInt, _ := strconv.Atoi(dataSplit[1])
			handList = append(handList, dataSplit[0])
			betList[dataSplit[0]] = betInt
		}
		// sort.Strings(handList)

		sort.Sort(sort.Reverse(sort.StringSlice(handList)))
		// sort.SliceStable(handList, func(i, j int) bool {
		//     return checkBest(ind, handList[i], handList[j])
		// })
		sort.Slice(handList, func(i, j int) bool {
			// Get the two strings to compare
			strI := handList[i]
			strJ := handList[j]

			// Find the minimum length to compare character by character
			minLen := 5
			// Compare character by character based on custom order
			for k := 0; k < minLen; k++ {
				// Convert characters to strings for map lookup
				charI := string(strI[k])
				charJ := string(strJ[k])

				// Lookup the order value for each character
				orderI := ind[charI]
				orderJ := ind[charJ]

				// Compare based on the custom order
				if orderI != orderJ {
					return orderI > orderJ
				}
			}

			// If all compared characters are equal, the shorter string comes first
			return len(strI) < len(strJ)
		})

		resPart := 0
		for _, value := range handList {
			resPart = c * betList[value]
			c += 1
			result += resPart
		}
	}
	fmt.Println(result)
	return 0
}


func partTwo(data []string) int {

	ind_joker := map[string]int{"A": 1, "K": 2, "Q": 3, "T": 4, "9": 5, "8": 6, "7": 7, "6": 8, "5": 9, "4": 10, "3": 11, "2": 12, "J": 13}

	fmt.Println(data)
	results := make(map[int][]string)
	result := 0
	c := 1

	for _, data := range data {
		dataSplit := strings.Split(data, " ")
		bets := checkStringJoker(dataSplit[0])
        
		betLen := len(bets)
		v := make([]int, 0, betLen)

		for _, value := range bets {
			v = append(v, value)
		}
		if betLen == 1 {
			results[1] = slices.Insert(results[1], len(results[1]), data)
		} else if betLen == 2 {
			containNum := slices.Contains(v, 1)
			if containNum == false {
				results[3] = slices.Insert(results[3], len(results[3]), data)
			} else {
				results[2] = slices.Insert(results[2], len(results[2]), data)
			}
		} else if betLen == 3 {
			containNum2 := slices.Contains(v, 2)
			containNum3 := slices.Contains(v, 3)
			if containNum3 == true {
				results[4] = slices.Insert(results[4], len(results[4]), data)
			} else if containNum2 == true {
				results[5] = slices.Insert(results[5], len(results[5]), data)
			} else {
				results[6] = slices.Insert(results[6], len(results[6]), data)
			}

		} else if betLen == 4 {
			results[7] = slices.Insert(results[7], len(results[7]), data)
		} else {
			results[8] = slices.Insert(results[8], len(results[8]), data)
		}

		fmt.Println(results)
	}
	keys := make([]int, 0, len(results))

	for k := range results {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, i := range keys {
		handList := make([]string, 0, len(results[i]))
		betList := make(map[string]int)
		for _, value := range results[i] {
			dataSplit := strings.Split(value, " ")
			betInt, _ := strconv.Atoi(dataSplit[1])
			handList = append(handList, dataSplit[0])
			betList[dataSplit[0]] = betInt
		}
		// sort.Strings(handList)

		sort.Sort(sort.Reverse(sort.StringSlice(handList)))
		// sort.SliceStable(handList, func(i, j int) bool {
		//     return checkBest(ind, handList[i], handList[j])
		// })
		sort.Slice(handList, func(i, j int) bool {
			// Get the two strings to compare
			strI := handList[i]
			strJ := handList[j]

			// Find the minimum length to compare character by character
			minLen := 5
			// Compare character by character based on custom order
			for k := 0; k < minLen; k++ {
				// Convert characters to strings for map lookup
				charI := string(strI[k])
				charJ := string(strJ[k])

				// Lookup the order value for each character
				orderI := ind_joker[charI]
				orderJ := ind_joker[charJ]

				// Compare based on the custom order
				if orderI != orderJ {
					return orderI > orderJ
				}
			}

			// If all compared characters are equal, the shorter string comes first
			return len(strI) < len(strJ)
		})

		resPart := 0
		for _, value := range handList {
			fmt.Println(c, value)
			resPart = c * betList[value]
			c += 1
			result += resPart
		}
	}
	fmt.Println(result)
	return 0
}
