package main

import (
    "bufio"
    "fmt"
    "log"
	"regexp"
    "os"
	"strconv"
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

	result := 0
	re := regexp.MustCompile("[1-9][0-9]*|one|two|three|four|five|six|seven|eight|nine")
	toDigit := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()
		numeric := re.FindFirstSting(line, -1)
		if len(numeric) == 0 {
			fmt.Println(line)
			continue
		}

		last, ok := toDigit[numeric[len(numeric) - 1]]
		if !ok {
			last, _ = strconv.Atoi(numeric[len(numeric) - 1])	
			last = findRightDigit(last)
		}
		
		first, ok := toDigit[numeric[0]]
		if !ok {
			first, _ = strconv.Atoi(numeric[0])
			first = findLeftDigit(first)
		}		

		num := 10 * first + last
		fmt.Println(line, numeric, num)
		result += num
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println(result)
}

func findRightDigit(i int) int {
	return i % 10
}

func findLeftDigit(i int) int {	
	for i >= 10 {
		i = i / 10
	}
	return i
}