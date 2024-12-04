package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	totalSafe := 0
	totalSafeTol := 0
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		row := ToInt(strings.Fields(line))
		if IsSafe(row, false) {
			totalSafe += 1
		}
		if IsSafe(row, true) {
			totalSafeTol += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part1:", totalSafe)
	fmt.Println("Part2:", totalSafeTol)
}
func ToInt(row []string) []int {
	var result []int
	for _, s := range row {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}

func IsSafe(row []int, tolerate bool) bool {
	decreasing := false
	if row[0] > row[len(row)-1] {
		decreasing = true
	}

	var prev int
	for i, v := range row {
		if i == 0 {
			prev = v
		} else {
			diff := v - prev
			if decreasing {
				diff *= -1
			}
			if diff < 1 || diff > 3 {
				if tolerate {
					return makeSubProblems(row)
				}
				return false
			}
			prev = v
		}
	}
	return true
}

func makeSubProblems(row []int) bool {
	fmt.Println()
	fmt.Println(row, len(row))
	for i := range len(row) {
		sub := []int{}
		sub = append(sub, row[:i]...)
		sub = append(sub, row[i+1:]...)
		if IsSafe(sub, false) {
			return true
		}
	}
	fmt.Println(row, len(row))
	return false
}
