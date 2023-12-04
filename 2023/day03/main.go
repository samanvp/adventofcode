package main

import (
    "bufio"
    "fmt"
    "log"
    "strconv"
    "os"
)

type Pos struct {
    line int
    col int
}

func main() {
    digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
    symbols := make(map[int][]int)

    gears := make(map[int][]int)
    ratios := make(map[Pos][]int)
    
    // open file
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    lineNo := 0
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()
		for i, c := range line {
			if !contains(digits, c) && c != '.' {
				a, ok := symbols[lineNo]
                if !ok {
                    symbols[lineNo] = []int{i}
                } else {
                    symbols[lineNo] = append(a, i)
                }
			}

            if c == '*' {
                a, ok := gears[lineNo]
                if !ok {
                    gears[lineNo] = []int{i}
                } else {
                    gears[lineNo] = append(a, i)
                }
            }
		}
        lineNo += 1
    }
    //fmt.Println(symbols)
    f.Seek(0, 0)
    // read the file line by line using scanner
    scanner = bufio.NewScanner(f)

    p1Result := 0
    lineNo = 0
    for scanner.Scan() {
        line := scanner.Text()

        var num string
        start := -1
        for i, c := range line {
            if contains(digits, c) {
                if start == -1 {
                    start = i
                    num = string(c)
                } else {
                    num = num + string(c)
                }
            } else if start != -1 {
                value, _ := strconv.Atoi(num)
                if hasAdjacent(lineNo, start, len(num), symbols) != nil {
                    p1Result += value
                }

                p := hasAdjacent(lineNo, start, len(num), gears) 
                if p != nil {
                    a, ok := ratios[*p]
                    if !ok {
                        ratios[*p] = []int{value}
                    } else {
                        ratios[*p] = append(a, value)
                    }
                }
                start = -1
            }
        }
        if start != -1 {
            value, _ := strconv.Atoi(num)
            if hasAdjacent(lineNo, start, len(num), symbols) != nil {
                p1Result += value
            }

            p := hasAdjacent(lineNo, start, len(num), gears) 
            if p != nil {
                a, ok := ratios[*p]
                if !ok {
                    ratios[*p] = []int{value}
                } else {
                    ratios[*p] = append(a, value)
                }
            }
        }
        lineNo += 1
    }
    fmt.Println("Part1: ", p1Result)

    p2Result := 0
    for _, v := range ratios {
        if len(v) == 2 {
            p2Result += v[0] * v[1]
        }
    }
    fmt.Println("Part2: ", p2Result)
}

func contains(s []rune, e rune) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func hasAdjacent(lineNo, start, len int, symbols map[int][]int) *Pos {
    for l := lineNo - 1; l <= lineNo + 1; l++ {
        s, ok := symbols[l]
        if ok {
            for _, v := range s {
                if v >= start - 1 && v <= start + len {
                    return &Pos{line: l, col: v}
                }
            }
        }
    }
    return nil
}
