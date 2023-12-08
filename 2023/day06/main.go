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
    var times, distances []int
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()
        if strings.HasPrefix(line, "Time:") {
            times = GetInts(strings.Split(line, ":")[1])
        }
        if strings.HasPrefix(line, "Distance:") {
            distances = GetInts(strings.Split(line, ":")[1])
        }
    }

    fmt.Println(times)
    fmt.Println(distances)

    p1Result := 1
    for i := 0; i < len(times); i++ {
        v := FindNumWins(times[i], distances[i])
        p1Result *= v
    }
    fmt.Println(p1Result)

    time := ""
    distance := ""
    for i := 0; i < len(times); i++ {
        time = time + strconv.Itoa(times[i])
        distance = distance + strconv.Itoa(distances[i])
    }

    timeI, _ := strconv.Atoi(time)
    distanceI, _ := strconv.Atoi(distance)
    p2Result := FindNumWins(timeI, distanceI)
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

func FindNumWins(t, d int) int {
    count := 0
    /*
    m := t / 2
    for m * (t - m) > d {
        count += 1
        m -= 1
    }
    return (count - (t + 1) % 2) * 2
    */
    for x := 1; x < t; x ++ {
        if x * (t - x) > d {
            count += 1
        }
    }
    return count
}