package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "regexp"
	// "slices"
	"strconv"
)

type Races struct {
    time int
    distance int
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
	fmt.Println("Advent of Code 2023: Day 6")
	fmt.Println("===========================")

	data := readFile("2023/data/day_6.txt")

	// Part 1
	fmt.Println("\tPart 1")
	result_1 := partOne(data)
	fmt.Printf("-> %d\n", result_1)
	// Part 2
	fmt.Println("\tPart 2")
	result_2 := partTwo(data)
	fmt.Printf("-> %d\n", result_2)
}

func getTimeDistance(data []string) []Races{
    var races []Races
    r := regexp.MustCompile(`[0-9]+`)

    for i, line := range data {
        lineSplit := strings.Trim(strings.Split(line, ":")[1], " ")
        
        for j, v := range r.FindAllString(lineSplit, -1) {
            valInt, _ := strconv.Atoi(strings.Trim(v, " "))
            if i == 0 {
                race := Races{time: valInt}
                races = append(races, race)
            } else {
                races[j].distance = valInt
            }
        }

    }

    return races

}


func getSingleTimeDistance(data []string) Races{
    var race Races

    for i, line := range data {
        lineSplit := strings.Trim(strings.Split(line, ":")[1], " ")

        lineSplit = strings.ReplaceAll(lineSplit, " ", "")
        
        valInt, _ := strconv.Atoi(strings.Trim(lineSplit, " "))
        if i == 0 {
            race.time = valInt
        } else {
            race.distance = valInt
        }

    }

    return race

}


func partOne(data []string) int {

    races := getTimeDistance(data)
    fmt.Println(races)
    result := 1
    for _, race := range races {
        fmt.Println(race)

        c := 0
        for i := 1; i<race.time; i++ {
            runTime := race.time - i
            distanceCovered := runTime * i
            // fmt.Println(distanceCovered)

            if distanceCovered > race.distance {
                c+=1
                // fmt.Println(distanceCovered)
            }
        }
        result*=c

    }
    fmt.Println("---->",result)
    return result
}


func partTwo(data []string) int {

    race := getSingleTimeDistance(data)
    result := 0
    fmt.Println(race)

    for i := 1; i<race.time; i++ {
        runTime := race.time - i
        distanceCovered := runTime * i
        // fmt.Println(distanceCovered)

        if distanceCovered > race.distance {
            result+=1
            // fmt.Println(distanceCovered)
        }
    }

    fmt.Println("---->",result)
    return result
}
