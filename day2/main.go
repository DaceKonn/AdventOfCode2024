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
	fmt.Println("\n# Evaluate safety")
	fmt.Println("\n# Sum reports")
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
