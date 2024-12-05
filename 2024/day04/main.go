package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	input := make([][]byte, 0, 0)
	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		input = append(input, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("num rows:", len(input))
	fmt.Println("num columns:", len(input[0]))

	totalCount := 0
	for i := 0; i < len(input); i += 1 {
		for j := 0; j < len(input[0]); j += 1 {
			totalCount += FindXmas(input, i, j, []byte("XMAS"))
		}
	}
	fmt.Println("Part1:", totalCount)
	fmt.Println("Part2:", findCrossMas(input))

}

func FindXmas(input [][]byte, i int, j int, xmas []byte) int {
	count := 0
	for x := -1; x <= 1; x += 1 {
		for y := -1; y <= 1; y += 1 {
			if x == 0 && y == 0 {
				continue
			}
			found := true
			for w := 0; w < len(xmas); w++ {
				newI := i + x*w
				newJ := j + y*w
				if newI < 0 || newJ < 0 || newI >= len(input) || newJ >= len(input[0]) {
					found = false
					break
				}
				if input[newI][newJ] != xmas[w] {
					found = false
					break
				}
			}
			if found {
				count += 1
			}
		}
	}
	return count
}

func findCrossMas(input [][]byte) int {
	count := 0
	for i := 1; i < len(input)-1; i += 1 {
		for j := 1; j < len(input[0])-1; j += 1 {
			if input[i][j] != byte('A') {
				continue
			}

			NW := input[i-1][j-1]
			NE := input[i-1][j+1]
			SW := input[i+1][j-1]
			SE := input[i+1][j+1]

			M := byte('M')
			S := byte('S')
			if ((NW == M && SE == S) || (NW == S && SE == M)) && ((NE == M && SW == S) || (NE == S && SW == M)) {
				count += 1
			}
		}
	}
	return count
}
