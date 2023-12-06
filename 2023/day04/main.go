package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strings"
)

const (
	INPUT = "input.txt"
	INPUT_LEN = 198
	//INPUT = "test.txt"
	//INPUT_LEN = 6
)

func main() {
    // open file
    f, err := os.Open(INPUT)
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

	copies := make([]int, INPUT_LEN)
	for i := 0; i < INPUT_LEN; i++ {
		copies[i] = 1
	}
	p1Result := 0
	lineNo := 0 
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()

		card := strings.Split(line, ":")

		numbers := strings.Split(card[1], "|")

		winings := strings.Split(strings.TrimSpace(numbers[0]), " ")
		haves := strings.Split(strings.TrimSpace(numbers[1]), " ")

		wins := 0
		score := 0
		for i, h := range haves {
			if len(strings.TrimSpace(h)) == 0 {
				continue
			}
			if contains(haves[i+1:], h) {
				fmt.Println(h)
				fmt.Println(line)
				continue
			}
			if contains(winings, h) {
				wins += 1
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		p1Result += score

		fmt.Println(wins)
		for i := lineNo + 1; (i <= lineNo + wins) && (i < INPUT_LEN); i ++ {
			copies[i] += copies[lineNo]
			fmt.Println(copies)
		}
		lineNo += 1
    }

	fmt.Println(p1Result)

	p2Result := 0
	for i := 0; i < INPUT_LEN; i++ {
		p2Result += copies[i]
	}
	fmt.Println(p2Result)
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if strings.TrimSpace(a) == strings.TrimSpace(e) {
            return true
        }
    }
    return false
}


