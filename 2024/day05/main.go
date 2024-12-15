package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	rules := GetRules(scanner)
	fmt.Println("Rules are loaded...")

	total1 := 0
	total2 := 0
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		midValue := IsPageOrdered(line, rules)
		if midValue != -1 {
			total1 += midValue
		} else {
			total2 += OrderPage(line, rules)
		}
	}

	fmt.Println("Part1: ", total1)
	fmt.Println("Part2: ", total2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetRules(scanner *bufio.Scanner) map[int][]int {
	rules := make(map[int][]int)
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		row := strings.Split(line, "|")
		first, _ := strconv.Atoi(row[0])
		second, _ := strconv.Atoi(row[1])

		arr, ok := rules[first]
		if ok {
			rules[first] = append(arr, second)
		} else {
			rules[first] = []int{second}
		}
	}

	// sort slices
	for _, value := range rules {
		sort.Ints(value)
	}
	return rules
}

func OrderPage(line string, rules map[int][]int) int {
	input := ToInt(line)

	reset := false
	for s := len(input) - 1; s > 0; s -= 1 {
		if reset {
			s = len(input) - 1
			reset = false
		}
		for f := s - 1; f >= 0; f -= 1 {
			first := input[f]
			second := input[s]
			afters, ok := rules[second]
			if !ok {
				break
			}
			if SearchArray(afters, first) {
				input[s] = first
				input[f] = second
				reset = true
				break
			}
		}
	}

	return input[len(input)/2]
}

func IsPageOrdered(line string, rules map[int][]int) int {
	input := ToInt(line)

	for s := len(input) - 1; s > 0; s -= 1 {
		for f := s - 1; f >= 0; f -= 1 {
			first := input[f]
			second := input[s]
			afters, ok := rules[second]
			if !ok {
				break
			}
			if SearchArray(afters, first) {
				return -1
			}
		}
	}

	return input[len(input)/2]
}

func SearchArray(a []int, x int) bool {
	index := sort.SearchInts(a, x)
	if index == len(a) || a[index] != x {
		return false
	}
	return true
}

func ToInt(line string) []int {
	input := strings.Split(line, ",")
	output := make([]int, 0)
	for _, s := range input {
		i, _ := strconv.Atoi(s)
		output = append(output, i)
	}

	return output
}
