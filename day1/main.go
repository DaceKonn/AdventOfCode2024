package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const file string = "input_real.txt"

func main() {
	fmt.Println("# Advent of Code 2024")

	fmt.Println("\n# Read input file")
	leftList, rightList, err := readFile(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n## Lists read")

	fmt.Println("\n\tLeft list:")
	logList(leftList)
	fmt.Println("\n\tRight list:")
	logList(rightList)

	fmt.Println("\n# Process lists")
	//	fmt.Println("\n## Removing duplicates")
	//	fmt.Printf("\tLeft list size: %d\n\tRight list size: %d\n", len(leftList), len(rightList))
	//	uniqueLeftList := removeDuplicates(leftList)
	//	uniqueRightList := removeDuplicates(rightList)
	//	fmt.Printf("\tLeft list size: %d\n\tRight list size: %d\n", len(uniqueLeftList), len(uniqueRightList))
	uniqueLeftList := leftList
	uniqueRightList := rightList

	fmt.Println("\n## Sorting lists")
	sort.Ints(uniqueLeftList)
	sort.Ints(uniqueRightList)

	fmt.Println("\n\tLeft list:")
	logList(uniqueLeftList)
	fmt.Println("\n\tRight list:")
	logList(uniqueRightList)

	fmt.Println("\n# Calculate distance")
	fmt.Println("## Calculate partial distances")
	distances, err := calculateDistances(uniqueLeftList, uniqueRightList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\tDistances: %d\n", distances)

	fmt.Println("\n## Calculate distance")
	var distance int = 0
	for _, value := range distances {
		distance += value
	}
	fmt.Printf("\tDistance is: %d", distance)

	fmt.Println("\n# Find similarity score")
	apperancesInRight := getNumberOfApperances(uniqueRightList)
	fmt.Printf("\tapperances in right list: %d\n", apperancesInRight)
	similarityIndexes := calculateSimilarityIndexes(uniqueLeftList, apperancesInRight)
	fmt.Printf("\tSimilarity indexes: %d\n", similarityIndexes)
	var similarityScore int = 0
	for _, value := range similarityIndexes {
		similarityScore += value
	}
	fmt.Printf("\tSimilarity score is: %d", similarityScore)
}

func calculateSimilarityIndexes(uniqueLeftList []int, apperancesInRight map[int]int) (result []int) {
	for _, value := range uniqueLeftList {
		result = append(result, value*apperancesInRight[value])
	}
	return result
}

func getNumberOfApperances(uniqueRightList []int) map[int]int {
	result := make(map[int]int)
	for _, value := range uniqueRightList {
		result[value] += 1
	}

	return result
}

func calculateDistances(uniqueLeftList, uniqueRightList []int) (distances []int, err error) {
	if len(uniqueLeftList) != len(uniqueRightList) {
		return nil, errors.New("Length of left and right lists are not even")
	}

	sharedLength := len(uniqueLeftList)
	for i := 0; i < sharedLength; i++ {
		distance := uniqueLeftList[i] - uniqueRightList[i]
		if distance < 0 {
			distance = -distance
		}
		distances = append(distances, distance)
	}

	return distances, nil
}

func readFile(name string) (leftList []int, rightList []int, err error) {
	file, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
		return nil, nil, errors.New("File open error")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("\tReading line: \"%s\"\n", line)
		parts := strings.Split(line, "   ")
		leftInt := parseInt(parts[0])
		rightInt := parseInt(parts[1])
		fmt.Printf("\t\tParsed left int: [%d] \n\t\tParsed right int: [%d]\n", leftInt, rightInt)

		leftList = append(leftList, leftInt)
		rightList = append(rightList, rightInt)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, nil, errors.New("Scanner error")
	}

	return leftList, rightList, nil
}

func parseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return num
}

func logList(list []int) {
	for _, value := range list {
		fmt.Printf("\t%d\n", value)
	}
}

func removeDuplicates(list []int) (result []int) {
	unique := make(map[int]bool)
	for _, value := range list {
		if !unique[value] {
			unique[value] = true
			result = append(result, value)
		}
	}
	return result
}
