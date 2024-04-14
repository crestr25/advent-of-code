package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Ranges struct {
	start int
	end   int
}

type Maps struct {
	sourceStart      []int
	destinationStart []int
	step             []int
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
	fmt.Println("Advent of Code 2023: Day 5")
	fmt.Println("===========================")

	data := readFile("2023/data/day_5.txt")

	// Part 1
	fmt.Println("\tPart 1")
	result_1 := partOne(data)
	fmt.Printf("-> %d\n", result_1)
	// Part 2
	fmt.Println("\tPart 2")
	result_2 := partTwo(data)
	fmt.Printf("-> %d\n", result_2)
}

func getSeeds(row string) []int {
	seedsStr := strings.Split(row, ": ")[1]
	seedsStrList := strings.Split(string(seedsStr), " ")
	var iSeeds []int

	for _, seed := range seedsStrList {
		seedInt, _ := strconv.Atoi(seed)
		iSeeds = append(iSeeds, seedInt)
	}
	return iSeeds
}

func getMaps(data []string) map[string]*Maps {

	maps := make(map[string]*Maps)
	var currentMapKey string

	for _, row := range data[2:] {
		if row == "" {
			continue
		}
		rowStr := strings.Split(row, " ")[0]
		_, err := strconv.Atoi(rowStr)
		if err != nil {
			currentMapKey = rowStr
			maps[currentMapKey] = new(Maps)
		} else {
			seedsStr := strings.Split(row, " ")
			var mapVals []int
			for _, val := range seedsStr {
				valInt, _ := strconv.Atoi(val)
				mapVals = append(mapVals, valInt)
			}
			maps[currentMapKey].sourceStart = append(maps[currentMapKey].sourceStart, mapVals[1])
			maps[currentMapKey].destinationStart = append(maps[currentMapKey].destinationStart, mapVals[0])
			maps[currentMapKey].step = append(maps[currentMapKey].step, mapVals[2])
		}
	}

	return maps

}

func seedRange(seed []int, sourceRange, destRange, step []int) []int {

	newSeed := make([]int, len(seed))
	for i := range newSeed {
		newSeed[i] = -1
	}

	for i := 0; i < len(sourceRange); i++ {
		for j, s := range seed {
			if s >= sourceRange[i] && s < (sourceRange[i]+step[i]) {
				convSeed := destRange[i] + (s - sourceRange[i])
				// fmt.Println(seed, convSeed, i)
				newSeed[j] = convSeed
			}

		}

	}

	for i, s := range newSeed {
		if s == -1 {
			newSeed[i] = seed[i]
		}
	}

	return newSeed

}

func partOne(data []string) int {

	iSeeds := getSeeds(data[0])

	maps := getMaps(data)
	// fmt.Println(maps)
	for _, row := range data[1:] {
		rowStr := strings.Split(row, " ")[0]
		_, err := strconv.Atoi(rowStr)
		if err == nil || row == "" {
			continue
		}
		// fmt.Println(maps[/* ro */wStr])
		iSeeds = seedRange(iSeeds, maps[rowStr].sourceStart, maps[rowStr].destinationStart, maps[rowStr].step)
	}
	return slices.Min(iSeeds)
}

func partTwo(data []string) int {

	iSeeds := getSeeds(data[0])

	iSeedsRange := seedRanges(iSeeds)
	maps := getMaps(data)
	var newSeedsRange []Ranges

	for _, row := range data[1:] {
		rowStr := strings.Split(row, " ")[0]
		_, err := strconv.Atoi(rowStr)
		if err == nil || row == "" {
			continue
		}

		newSeedsRange = []Ranges{}
		for len(iSeedsRange) > 0 {

			notMap := true
			// get the current seed Range
			currentSeed := iSeedsRange[0]
			// pop from slice
			iSeedsRange = iSeedsRange[1:]
            // iterate through all the mappins in a block for the current seed
			for i := 0; i < len(maps[rowStr].destinationStart); i++ {
                // Get the mapping values
				sourceMap := maps[rowStr].sourceStart[i]
				destMap := maps[rowStr].destinationStart[i]
				step := maps[rowStr].step[i]
				// get lower bound
				minBound := max(sourceMap, currentSeed.start)
				// get upper bound
				upBound := min(sourceMap+step, currentSeed.end)
                // range exists
				if minBound < upBound {
					notMap = false
					newSeedsRange = append(newSeedsRange, Ranges{
						start: minBound - sourceMap + destMap,
						end:   upBound - sourceMap + destMap})
                    // create new leftover to right range if exists
					if minBound > currentSeed.start {
						iSeedsRange = append(iSeedsRange, Ranges{
							start: currentSeed.start,
							end:   minBound,
						})
					}
                    // create new leftover to right range if exists
					if currentSeed.end > upBound {
						iSeedsRange = append(iSeedsRange, Ranges{
							start: upBound,
							end:   currentSeed.end,
						})
					}
					break

				}
			}

			if notMap == true {
                // no mapping applied range as is
				newSeedsRange = append(newSeedsRange, currentSeed)
			}
		}
        // all new ranges are the new input to continue
		iSeedsRange = newSeedsRange
	}

    // get the minimun
	var minVals []int
	for _, seedRange := range newSeedsRange {
		minVals = append(minVals, seedRange.start)
	}
	return slices.Min(minVals)
}

func seedRanges(iSeeds []int) []Ranges {

	var seedsRange []Ranges

	for i := 0; i < len(iSeeds)-1; i = i + 2 {
		seedRange := Ranges{start: iSeeds[i], end: iSeeds[i] + iSeeds[i+1]}
		seedsRange = append(seedsRange, seedRange)
	}

	return seedsRange
}
