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

	result := 0
	power := 0
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()
		game := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Split(game[0], " ")[1])

		invalid := false
		red := 0
		green := 0
		blue := 0
		sets := strings.Split(game[1], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				values := strings.Split(strings.TrimSpace(cube), " ")

				count, _ := strconv.Atoi(values[0])
				color := strings.TrimSpace(values[1])
				if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
					invalid = true
				}
				if color == "red" {
					red = max(red, count)
				}
				if color == "green" {
					green = max(green, count)
				}
				if color == "blue" {
					blue = max(blue, count)
				}
			}
		}
		power += red * green * blue
		if !invalid {
			result += id
		}

    }
	fmt.Println(result)
	fmt.Println(power)
}


func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}