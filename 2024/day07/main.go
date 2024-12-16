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

	total1 := 0
	total2 := 0
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		row := strings.Split(line, ":")
		wanted, _ := strconv.Atoi(row[0])
		operands := ToInt(row[1])
		if InsertOperators(0, operands, wanted) {
			total1 += wanted
		}

		if InsertOperators2(0, operands, wanted) {
			total2 += wanted
		}
	}

	fmt.Println("Part1: ", total1)
	fmt.Println("Part2: ", total2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func InsertOperators2(current int, operands []int, wanted int) bool {
	if current == wanted && len(operands) == 0 {
		return true
	}

	if current > wanted || len(operands) == 0 {
		return false
	}

	if current == 0 {
		return InsertOperators2(operands[0], operands[1:], wanted)
	}

	next := operands[0]
	if InsertOperators2(current+next, operands[1:], wanted) {
		return true
	}

	if InsertOperators2(current*next, operands[1:], wanted) {
		return true
	}

	newCurrent, _ := strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(next))
	if InsertOperators2(newCurrent, operands[1:], wanted) {
		return true
	}
	return false
}

func InsertOperators(current int, operands []int, wanted int) bool {
	if current == wanted && len(operands) == 0 {
		return true
	}

	if current > wanted || len(operands) == 0 {
		return false
	}

	if current == 0 {
		return InsertOperators(operands[0], operands[1:], wanted)
	}

	next := operands[0]
	return InsertOperators(current+next, operands[1:], wanted) || InsertOperators(current*next, operands[1:], wanted)
}

func ToInt(line string) []int {
	input := strings.Fields(line)
	output := make([]int, 0)
	for _, s := range input {
		i, _ := strconv.Atoi(s)
		output = append(output, i)
	}

	return output
}
