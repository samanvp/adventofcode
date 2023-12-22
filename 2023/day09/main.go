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
    p1Result := 0
    p2Result := 0
    var seq []int 
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()
        seq = GetInts(line)
        p1Result += FindNext(seq)
        p2Result += FindPrev(seq)
    }

    fmt.Println(p1Result)
    fmt.Println(p2Result)
}

func GetInts(s string) []int {
	nums := strings.Fields(s)
	var result []int
	for _, num := range nums {
		v, _ := strconv.Atoi(num)
		result = append(result, v)
	}
	return result
}

func FindNext(seq []int) int {
    if IsAllZeros(seq) {
        return 0
    }

    var diffs []int
    for i := 1; i < len(seq); i ++ {
        diffs = append(diffs, seq[i] - seq[i-1])
    }
    return FindNext(diffs) + seq[len(seq) - 1]
}

func FindPrev(seq []int) int {
    if IsAllZeros(seq) {
        return 0
    }

    var diffs []int
    for i := 1; i < len(seq); i ++ {
        diffs = append(diffs, seq[i] - seq[i-1])
    }
    return  seq[0] - FindPrev(diffs)
}


func IsAllZeros(seq []int) bool {
    for _, n := range seq {
        if n != 0 {
            return false
        }
    }
    return true
}