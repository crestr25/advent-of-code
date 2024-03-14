package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	fmt.Println("Advent of Code 2023: Day 3")
	fmt.Println("===========================")

	data := readFile("2023/data/day_3.txt")

	// Part 1
	fmt.Println("\tPart 1")
	result_1 := partOne(data)
	fmt.Printf("-> %d\n", result_1)
	// Part 2
	fmt.Println("\tPart 2")
	result_2 := partTwo(data)
	fmt.Printf("-> %d\n", result_2)
}

// func checkAround(num string, prev_line string, line string, next_line string) bool {
func checkLeft(left int, data string, rePattern *regexp.Regexp) bool {
    if left < 0{
        return false
    }
    match := rePattern.FindString(string(data[left]))

    if match == "" {
        return false
    }

    return true

}

func checkRight(right int, data string, rePattern *regexp.Regexp) bool {
    if right >= len(data){
        return false
    }
    match := rePattern.FindString(string(data[right]))
    if match == "" {
        return false
    }
    return true
}
func checkUp(up int, left int, right int, data []string, rePattern *regexp.Regexp) bool {
    if up < 0 {
        return false
    }

    if left < 0 {
        left = 0
    }

    match := ""

    if right >= len(data[up]){
        right = len(data[up])
        match = rePattern.FindString(string(data[up][left:right]))
    } else {
        match = rePattern.FindString(string(data[up][left:right + 1]))
    }

    if match == "" {
        return false
    }

    return true
}
func checkDown(down int, left int, right int, data []string, rePattern *regexp.Regexp) bool {
    if down >= len(data) {
        return false
    }

    if left < 0 {
        left = 0
    }

    match := ""

    if right >= len(data[down]){
        right = len(data[down])
        match = rePattern.FindString(string(data[down][left:right]))
    } else {
        match = rePattern.FindString(string(data[down][left:right + 1]))
    }
    if match == "" {
        return false
    }

    return true
}

func checkAround(ind []int, i int, data []string) bool {

    // Compile both regex
    reEspChar := regexp.MustCompile("[^a-zA-z0-9_.\n]")
    // reIndex := regexp.MustCompile(num)
    // find index for match
    // matchIndex := reIndex.FindStringIndex(data[i])
    up := i - 1
    down := i + 1
    left := ind[0] - 1 
    right := ind[1] 
    // fmt.Println(up, down, left, right)
    // Check left
    cLeft := checkLeft(left, data[i], reEspChar)
    // fmt.Println(cLeft)
    if cLeft == true {
        return true
    }
    // Check right
    cRight := checkRight(right, data[i], reEspChar)
    // fmt.Println(cRight)
    if cRight == true {
        return true
    }
    // Check up
    cUp := checkUp(up, left, right, data, reEspChar)
    // fmt.Println(cUp)
    if cUp == true {
        return true
    }
    // Check down
    cDown := checkDown(down, left, right, data, reEspChar)
    // fmt.Println(cDown)
    if cDown == true {
        return true
    }
   
     return false

}

func findIntersection(left int, right int, arr1 []int) bool {
	set := make(map[int]bool)

	// Create a set from the first array
    for i := left; i <= right; i++ {
		set[i] = true       // setting the initial value to true
	}

	// Check elements in the second array against the set
	for _, num := range arr1 {
		if set[num] {
            return true
		}
	}

	return false
}

func findNumber(left, right int, data string) []int {

    var res []int

    reNum := regexp.MustCompile("[0-9]+")
    indices := reNum.FindAllStringIndex(data, -1)
    for _, ind := range indices {
        fmt.Println(ind[0], ind[1] - 1, left, right)
        if findIntersection(left, right, []int{ind[0], ind[1] - 1}) == true {
            fmt.Println(ind[0], ind[1], left, right)
            num, err := strconv.Atoi(data[ind[0]:ind[1]])
            if err != nil {
                fmt.Println(err)
            }
            res = append(res, num)
        }

    }

    return res 
    
}

func checkAroundGear(ind []int, i int, data []string) int {

    // Compile both regex
    reNum := regexp.MustCompile("[0-9]")
    // reIndex := regexp.MustCompile(num)
    // find index for match
    // matchIndex := reIndex.FindStringIndex(data[i])
    var gearNums []int
    up := i - 1
    down := i + 1
    left := ind[0] - 1 
    right := ind[1] 
    fmt.Println(up, down, left, right)
    // Check left
    cLeft := checkLeft(left, data[i], reNum)
    fmt.Println(cLeft)
    if cLeft == true {
        numLeft := findNumber(left, left, data[i])        
        if len(numLeft) > 0 {
            gearNums = append(gearNums, numLeft...)
        }
    }
    // Check right
    cRight := checkRight(right, data[i], reNum)
    fmt.Println(cRight)
    if cRight == true {
        numRight := findNumber(right, right, data[i])        
        if len(numRight) > 0 {
            gearNums = append(gearNums, numRight...)
        }
    }
    // Check up
    cUp := checkUp(up, left, right, data, reNum)
    fmt.Println(cUp)
    if cUp == true {
        numUp := findNumber(left, right, data[up])        
        if len(numUp) > 0 {
            gearNums = append(gearNums, numUp...)
        }
    }
    // Check down
    cDown := checkDown(down, left, right, data, reNum)
    fmt.Println(cDown)
    if cDown == true {
        numDown := findNumber(left, right, data[down])        
        if len(numDown) > 0 {
            gearNums = append(gearNums, numDown...)
        }
        =BUSCARV(H5,$B$5:$B$1274,1,FALSO)
    }

    fmt.Println(gearNums)
    if len(gearNums) == 2 {
        res := 1
        // fmt.Println(gearNums)
        for _, i := range gearNums{
            res *= i
        }

        fmt.Println(res)
        return res
    }
   
     return 0

}

// 3..1*.3
// 3.1.1.3

func partOne(data []string) int {

    reNum := regexp.MustCompile("[0-9]+")
    result := 0

    for i := 0; i < len(data); i++ {
        // numMatches := reNum.FindAllString(data[i], -1)
        indices := reNum.FindAllStringIndex(data[i], -1)
        for _, ind := range indices {
            // fmt.Printf("-----> %d - ", data[i][ind[0]: ind[1]])
            // fmt.Println(ind[0])
            exist := checkAround(ind, i, data)

            if exist == true {
                // fmt.Println(data[ind[0]:ind[1]-1])
                iNum, _ := strconv.Atoi(string(data[i][ind[0]: ind[1]]))
                result += iNum
            }
        }

    }

    return result
}


func partTwo(data []string) int {

    reGear := regexp.MustCompile("[*]+")
    result := 0

    for i := 0; i < len(data); i++ {
        // numMatches := reNum.FindAllString(data[i], -1)
        indices := reGear.FindAllStringIndex(data[i], -1)
        for _, ind := range indices {
            fmt.Printf("-----> %d - ", ind)
            // fmt.Println(ind[0])
            num := checkAroundGear(ind, i, data)

            if num > 0 {
                // fmt.Println(data[ind[0]:ind[1]-1])
                // iNum, _ := strconv.Atoi(string(data[i][ind[0]: ind[1]]))
                result += num
            }
        }

    }

    return result
}
