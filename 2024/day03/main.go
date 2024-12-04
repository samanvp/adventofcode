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

	totalSum := 0
	doDontSum := 0
	enabled := true
	for scanner.Scan() {
		// do something with a line
		input := scanner.Text()
		totalSum += Mul(input)
		s, e := DoDont(input, enabled)
		doDontSum += s
		enabled = e
	}
	fmt.Println("Part1:", totalSum)
	fmt.Println("Part2:", doDontSum)
}

func findNum(s string) (int, int) {
	v, e := strconv.Atoi(s[:1])
	d := 1
	if e != nil {
		return -1, -1
	}
	if len(s) < 2 {
		return v, d
	}
	v2, e := strconv.Atoi(s[:2])
	if e == nil {
		v = v2
		d = 2
	}

	if len(s) < 3 {
		return v, d
	}
	v3, e := strconv.Atoi(s[:3])
	if e == nil {
		v = v3
		d = 3
	}

	if len(s) < 4 {
		return v, d
	}
	_, e = strconv.Atoi(s[:4])
	if e == nil {
		return -1, -1
	}
	return v, d
}

func findOperands(s string) (int, int) {
	v1, d1 := findNum(s)
	if v1 == -1 {
		return -1, -1
	}

	if s[d1] != byte(',') {
		return -1, -1
	}

	v2, d2 := findNum(s[(d1 + 1):])
	if v2 == -1 {
		return -1, -1
	}

	if s[d1+1+d2] != byte(')') {
		return -1, -1
	}
	return v1, v2
}

func Mul(input string) int {
	totalSum := 0
	for i := strings.Index(input, "mul("); i != -1; {
		v1, v2 := findOperands(input[i+4:])

		if v1 != -1 && v2 != -1 {
			totalSum += v1 * v2
		}
		input = input[i+1:]
		i = strings.Index(input, "mul(")
	}
	return totalSum
}

func DoDont(input string, enabled bool) (int, bool) {
	totalSum := 0
	for i := 0; i != -1; {
		if enabled {
			mul_i := strings.Index(input, "mul(")
			dont_i := strings.Index(input, "don't()")

			if mul_i == -1 && dont_i == -1 {
				break
			}
			if dont_i == -1 || mul_i < dont_i {
				i = mul_i
				v1, v2 := findOperands(input[i+4:])

				if v1 != -1 && v2 != -1 {
					totalSum += v1 * v2
				}
			} else {
				i = dont_i
				enabled = false
			}
		} else {
			do_i := strings.Index(input, "do()")
			if do_i > 0 {
				i = do_i
				enabled = true
			} else {
				break
			}
		}
		input = input[i+1:]
	}
	return totalSum, enabled
}
