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
    directions := ""
    network := make(map[string][]string)
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()

        if directions == "" {
            directions = strings.ReplaceAll(strings.ReplaceAll(line, "L", "0"), "R", "1")
            scanner.Scan()
        } else {
            AddLineToNetwork(network, line)
        }
    }

    /*
    p1Result := 0
    dIndex := 0
    current := "AAA"
    for current != "ZZZ" {
        d, _ := strconv.Atoi(string(directions[dIndex]))
        current = network[current][d]
        p1Result += 1
        dIndex = (dIndex + 1) % len(directions)
    }
    fmt.Println(p1Result)
    */

    state := []string{}
    for k , _ := range network {
        if strings.HasSuffix(k, "Z") {
            state = append(state, k)
        }
    }
    fmt.Println(state)

    
    fmt.Println(ReachToZ(state, directions, network))

}

func ReachToZ(state []string, directions string, network map[string][]string) int {
    result := 0
    dIndex := 0
    fmt.Println(state)
    for !AllHasSuffixZ(state) || result == 0 {
        d, _ := strconv.Atoi(string(directions[dIndex]))

        next := []string{}
        for _, s := range state {
            next = append(next, network[s][d])
        }
        state = next
        result += 1
        dIndex = (dIndex + 1) % len(directions)
        if result % 10000000 == 0 {
            fmt.Println(result)
        }
    }
    fmt.Println(state)
    return result
}

func AllHasSuffixZ(state []string) bool {
    for _, s := range state {
        if !strings.HasSuffix(s, "Z") {
            return false
        }
    }
    return true
}

func AddLineToNetwork(network map[string][]string, line string) {
    vals := strings.Split(line, " = ")
    key := vals[0]

    value := strings.Split(strings.Trim(vals[1], "()"), ", ")
    network[key] = value
}