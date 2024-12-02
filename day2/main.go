package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const file string = "input_test.txt"

func main() {
	fmt.Println("Advent of Code - day 2")
	fmt.Println("\n# Read file")

	reports, err := getReportsFromFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n## Reports read ")
	fmt.Printf("\t reports read: %d\n", len(reports))

	fmt.Println("\n# Calculate changes")
	var changes [][]int = calculateChanges(reports)
	fmt.Printf("\tChanges calculated: %d\n", len(changes))

	fmt.Println("\n# Evaluate safety")
	var safety map[int]bool = evaluateSafety(changes)
	fmt.Printf("\t%v\n", safety)

	fmt.Println("\n# Sum reports")
	var safetySum int = 0
	for _, value := range safety {
		if value {
			safetySum++
		}
	}
	fmt.Printf("\tNumber of safe reports: %d", safetySum)
}

func evaluateSafety(changes [][]int) map[int]bool {
	results := make(map[int]bool)

	for reportNo, changeRow := range changes {
		var increase bool
		var safe bool = true
		for changeIndx, change := range changeRow {
			if change > 3 || change < -3 || change == 0 {
				safe = false
				break
			}
			if changeIndx == 0 {
				increase = change > 0
			} else {
				newDirection := change > 0
				if newDirection != increase {
					safe = false
					break
				}
			}

		}
		results[reportNo] = safe
	}

	return results
}

func calculateChanges(reports [][]int) (changes [][]int) {
	for _, report := range reports {
		noOfLevels := len(report)
		noOfChanges := noOfLevels - 1
		var changesRow []int
		for i := 0; i < noOfChanges; i++ {
			changesRow = append(changesRow, report[i+1]-report[i])
		}
		fmt.Printf("\t%d\n", changesRow)
		changes = append(changes, changesRow)
	}
	return changes
}

func getReportsFromFile(fileName string) (reports [][]int, err error) {
	var file *os.File
	var inErr error
	file, inErr = os.Open(fileName)
	if inErr != nil {
		return reports, inErr
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var levels []int
		for _, value := range parts {
			levels = append(levels, parseInt(value))
		}
		reports = append(reports, levels)
	}
	return reports, nil
}

func parseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return num
}
