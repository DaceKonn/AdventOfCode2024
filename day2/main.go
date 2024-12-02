package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const file string = "input_real.txt"

type Level struct {
	Value        int
	Change       int
	BreaksSafety bool
}

func NewLevel(value int) Level {
	return Level{
		Value:        value,
		Change:       0,
		BreaksSafety: false,
	}
}

func (level Level) CopyAndClearLevel() Level {
	return NewLevel(level.Value)
}

type Report struct {
	Levels         []Level
	SafetyOverride bool
}

func NewReport(levels []Level) Report {
	return Report{
		Levels:         levels,
		SafetyOverride: false,
	}
}

func (report Report) CopyAndRemoveOneLevel(levelIndex int) Report {
	var newLevels []Level
	for index, level := range report.Levels {
		if index == levelIndex {
			continue
		}
		newLevels = append(newLevels, level.CopyAndClearLevel())
	}
	return NewReport(newLevels)
}

func (report Report) IsSafe() bool {
	if report.SafetyOverride {
		return true
	}
	for _, level := range report.Levels {
		if level.BreaksSafety {
			return false
		}
	}
	return true
}

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

	fmt.Printf("\n# Process reports\n")
	for reportNo, _ := range reports {
		fmt.Printf("\nProcessing report: %d\n", reportNo)
		processReport(&reports[reportNo], 1)
	}

	//	var changes [][]int = calculateChanges(reports)
	//	fmt.Printf("\tChanges calculated: %d\n", len(changes))
	//
	//	fmt.Printf("\t%v\n", safety)
	//
	//	fmt.Println("\n# Sum reports")
	//	var safetySum int = 0
	//	for _, value := range safety {
	//		if value {
	//			safetySum++
	//		}
	//	}
	//	fmt.Printf("\tNumber of safe reports: %d", safetySum)

	fmt.Printf("\n# Log all reports\n")
	for reportNo, report := range reports {
		fmt.Printf("\tReport: %d\tis safe:%v\twith override: %v\n", reportNo, report.IsSafe(), report.SafetyOverride)
		for _, level := range report.Levels {
			fmt.Printf("\t\tv: %d\tc: %d\tbs: %v\n", level.Value, level.Change, level.BreaksSafety)
		}
	}

	fmt.Println("\n# Result")
	var sumSafeReports int = 0
	for _, report := range reports {
		if report.IsSafe() {
			sumSafeReports++
		}
	}
	fmt.Printf("\nSafe reports: %d\n", sumSafeReports)
}

func processReport(report *Report, depth int) {
	fmt.Println("\n## Calculate changes")
	calculateChanges(report)

	fmt.Println("\n## Evaluate safety")
	evaluateSafety(report)

	logReport(report)

	if !report.IsSafe() && depth > 0 {
		fmt.Println("!! report not considered safe !!")
		//fmt.Println("!! processing report without elemnt at index 0")
		//newReport := report.CopyAndRemoveOneLevel(0)
		//processReport(&newReport, 0)
		//if !newReport.IsSafe() {
		fmt.Println("!! scanning for faulty Level")
		for index, _ := range report.Levels {
			fmt.Printf("!! processing report without element at index %d\n", index)
			newReport := report.CopyAndRemoveOneLevel(index)
			processReport(&newReport, 0)
			if !newReport.IsSafe() {
				fmt.Println("!! failed to clear report !!")
				continue
			} else {
				fmt.Println("+ + fixed report")
				report.SafetyOverride = true
				logReport(report)
				break
			}
		}
		//}
	} else if !report.IsSafe() && depth == 0 {
		fmt.Println("!! report still not safe rolling back recursion")
	}
}

func logReport(report *Report) {
	fmt.Println("\n## Log report")
	fmt.Printf("\tSafety Override: %v\tis safe: %v\n", report.SafetyOverride, report.IsSafe())
	for _, level := range report.Levels {
		fmt.Printf("\t\tv: %d\tc: %d\tbs: %v\n", level.Value, level.Change, level.BreaksSafety)
	}
}

func evaluateSafety(report *Report) {
	var increase bool
	for levelIndx, level := range report.Levels {
		change := level.Change
		if levelIndx == 0 {
			continue
		}
		if change > 3 || change < -3 || change == 0 {
			report.Levels[levelIndx].BreaksSafety = true
		}
		if levelIndx == 1 {
			increase = change > 0
		} else {
			newDirection := change > 0
			if newDirection != increase {
				report.Levels[levelIndx].BreaksSafety = true
				increase = newDirection
			}
		}

	}
}

func calculateChanges(report *Report) {
	noOfLevels := len(report.Levels)
	noOfChanges := noOfLevels - 1
	for i := 0; i < noOfChanges; i++ {
		report.Levels[i+1].Change = report.Levels[i+1].Value - report.Levels[i].Value
	}
}

func getReportsFromFile(fileName string) (reports []Report, err error) {
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
		var levels []Level
		for _, value := range parts {
			levels = append(levels, NewLevel(parseInt(value)))
		}
		reports = append(reports, NewReport(levels))
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
