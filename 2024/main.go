package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		row := strings.Fields(line)
		fmt.Println(row[0])
		fmt.Println(row[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
