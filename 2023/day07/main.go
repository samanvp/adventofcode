package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "slices"
	"strconv"
	"strings"
)

var scores map[rune]int
var scores2 map[rune]int

func main() {
    // open file
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    scores = map[rune]int{'A' : 14, 'K' : 13, 'Q' : 12, 'J' : 11, 'T' : 10, '9' : 9, '8' : 8, '7' : 7, '6' : 6, '5' : 5, '4' : 4, '3' : 3, '2' : 2}
    scores2 = map[rune]int{'A' : 14, 'K' : 13, 'Q' : 12, 'T' : 10, '9' : 9, '8' : 8, '7' : 7, '6' : 6, '5' : 5, '4' : 4, '3' : 3, '2' : 2, 'J' : 1}
    worths := []int{}
    bets := []int{}
   // read the file line by line using scanner
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()
        vals := strings.Fields(line)

        hand := vals[0]
        bet, _ := strconv.Atoi(vals[1])

        //worths = append(worths, GetWorth(hand))
        worths = append(worths, GetWorth2(hand))
        bets = append(bets, bet)
    }


    p1Result := 0
    rank := 1
    for len(worths) > 0 {
        min := slices.Min(worths)
        index := slices.Index(worths, min)
        
        p1Result += rank * bets[index]
   
        worths = slices.Delete(worths, index, index + 1)
        bets = slices.Delete(bets, index, index + 1)
        rank += 1
    }
    fmt.Println(p1Result)
}

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func GetWorth2(hand string) int {
    score := powInt(20, len(hand) + 1) * GetType2(hand)
    for i, r := range hand {
        score += powInt(20, len(hand) - i) * scores2[r]
    }
    return score
}

func GetType2(hand string) int {
    temp := make(map[rune]int)

    numJs := 0
    for _, c := range hand {
        if c == 'J' {
            numJs += 1
            continue
        }
        _, ok := temp[c]
        if ok {
            temp[c] = temp[c] + 1
        } else {
            temp[c] = 1
        }
    }

    counts := []int{0, 0, 0, 0, 0, 0}
    for _, v := range temp {
        counts[v] += 1
    }

    if counts[5] == 1  || (counts[4] == 1 && numJs == 1) || (counts[3] == 1 && numJs == 2) || (counts[2] == 1 && numJs == 3) || (counts[1] == 1 && numJs == 4) || numJs == 5{
        return 7
    }

    if counts[4] == 1 || (counts[3] == 1 && numJs == 1) || (counts[2] == 1 && numJs == 2) || (counts[1] == 2 && numJs == 3) {
        return 6
    }

    if (counts[3] == 1 && counts[2] == 1) || (counts[2] == 2 && numJs == 1) {
        return 5
    }

    if (counts[3] == 1 && counts[1] == 2) || (counts[2] == 1 && counts[1] == 2 && numJs == 1) || (counts[1] == 3 && numJs == 2) {
        return 4
    }

    if (counts[2] == 2 && counts[1] == 1) || (counts[2] == 1 && counts[1] == 2 && numJs == 1) {
        return 3
    }

    if (counts[2] == 1 && counts[1] == 3) || (counts[1] == 4 && numJs == 1) {
        return 2
    }

    if counts[1] == 5 {
        return 1
    }

    fmt.Println("Error", counts, numJs)
    return -1
}
func GetWorth(hand string) int {
    score := powInt(20, len(hand) + 1) * GetType(hand)
    for i, r := range hand {
        score += powInt(20, len(hand) - i) * scores[r]
    }
    return score
}

func GetType(hand string) int {
    temp := make(map[rune]int)

    for _, c := range hand {
        _, ok := temp[c]
        if ok {
            temp[c] = temp[c] + 1
        } else {
            temp[c] = 1
        }
    }

    counts := []int{0, 0, 0, 0, 0, 0}
    for _, v := range temp {
        counts[v] += 1
    }

    if counts[5] == 1 {
        return 7
    }

    if counts[4] == 1 {
        return 6
    }

    if counts[3] == 1 && counts[2] == 1 {
        return 5
    }

    if counts[3] == 1 && counts[1] == 2 {
        return 4
    }

    if counts[2] == 2 && counts[1] == 1 {
        return 3
    }

    if counts[2] == 1 && counts[1] == 3 {
        return 2
    }

    if counts[1] == 5 {
        return 1
    }

    fmt.Println("Error")
    return -1
}