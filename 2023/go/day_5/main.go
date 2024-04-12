package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "slices"
    "sync"
)

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
            maps[currentMapKey].destinationStart  = append(maps[currentMapKey].destinationStart, mapVals[0])
            maps[currentMapKey].step  = append(maps[currentMapKey].step, mapVals[2])
        }
    }

    return maps

}

func seedRange(seed []int, sourceRange, destRange, step []int) []int {

    newSeed := make([]int, len(seed))
    for i := range newSeed {
        newSeed[i] = -1
    }


    for i := 0; i < len(sourceRange);i++ {
        for j, s := range seed {
            if s >= sourceRange[i] && s < (sourceRange[i] + step[i]) {
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

func seedRangeReload(seed []int, sourceRange, destRange, step []int) []int {

    newSeed := make([]int, len(seed))
    for i := range newSeed {
        newSeed[i] = -1
    }


    for i := 0; i < len(sourceRange);i++ {
        for j, s := range seed {
            if s >= sourceRange[i] && s < (sourceRange[i] + step[i]) {
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

	fmt.Println(iSeeds)
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

    wg := sync.WaitGroup{}
    minRes := []int{}

    for i:=0;i<len(iSeeds);i=i+2{
        wg.Add(1)
        go func(data []string, seedStart, seedEnd int){
            num := seedBuildBF(data, seedStart, seedEnd) 
            minRes = append(minRes, num)
            wg.Done()
        }(data, iSeeds[i], iSeeds[i+1])
    }

    wg.Wait()
    fmt.Println(minRes)
	return slices.Min(minRes) 
    // return 0
}

func seedBuildBF(data []string, seedStart, seedEnd int) int {

    var iSeedsBF []int
    for j:=0;j<seedEnd;j++{
        iSeedsBF = append(iSeedsBF, seedStart + j)
    }

    maps := getMaps(data)
	// fmt.Println(maps)
    for _, row := range data[1:] {
        rowStr := strings.Split(row, " ")[0]
		_, err := strconv.Atoi(rowStr)
        if err == nil || row == "" {
            continue
        }
        iSeedsBF = seedRange(iSeedsBF, maps[rowStr].sourceStart, maps[rowStr].destinationStart, maps[rowStr].step)
	}

    return slices.Min(iSeedsBF)
}
