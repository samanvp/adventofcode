package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	start int
	end int
	length int
	dest int
}

type Ranges struct {
	ranges []*Range
}

func (rng Ranges) CategoryMap(source int) int {
	for _, r := range rng.ranges {
		//if source >= r.start && source < r.start + r.length {
		if source >= r.start && source < r.end {
			return r.dest + source - r.start
		}
	}
	return source
}

func CategoryMap(source int, ranges []*Range) int {
	for _, r := range ranges {
		//if source >= r.start && source < r.start + r.length {
		if source >= r.start && source < r.end {
			return r.dest + source - r.start
		}
	}
	return source
}

var s2s, s2f, f2w, w2l, l2t, t2h, h2l Ranges
func FindLocation(seed int) int {
	s := s2s.CategoryMap(seed)
	f := s2f.CategoryMap(s)
	w := f2w.CategoryMap(f)
	l := w2l.CategoryMap(w)
	t := l2t.CategoryMap(l)
	h := t2h.CategoryMap(t)
	return h2l.CategoryMap(h)
}
func ScannerToRanges(scanner *bufio.Scanner) Ranges {
	var result []*Range
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		result = append(result, LineToRange(line))
	}
	return Ranges{ranges: result}
}

/*
var s2s, s2f, f2w, w2l, l2t, t2h, h2l []*Range
func FindLocation(seed int) int {
	s := CategoryMap(seed, s2s)
	f := CategoryMap(s, s2f)
	w := CategoryMap(f, f2w)
	l := CategoryMap(w, w2l)
	t := CategoryMap(l, l2t)
	h := CategoryMap(t, t2h)
	return CategoryMap(h, h2l)
}
func ScannerToRanges(scanner *bufio.Scanner) []*Range {
	var result []*Range
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		result = append(result, LineToRange(line))
	}
	return result
}
*/

func LineToRange(line string) *Range {
	nums := strings.Split(strings.TrimSpace(line), " ")
	v1, _ := strconv.Atoi(nums[0])
	v2, _ := strconv.Atoi(nums[1])
	v3, _ := strconv.Atoi(nums[2])
	return &Range{start: v2, end: v2 + v3, length: v3, dest: v1}
}

func GetSeeds(s string) []int {
	nums := strings.Split(strings.TrimSpace(s), " ")
	var result []int
	for _, num := range nums {
		v, _ := strconv.Atoi(num)
		result = append(result, v)
	}
	return result
}

func main() {
    // open file
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

	var seeds []int
    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seeds = GetSeeds(strings.Split(line, ":")[1])
		}

		if strings.HasPrefix(line, "seed-to-soil map:") {
			s2s = ScannerToRanges(scanner)
		}

		if strings.HasPrefix(line, "soil-to-fertilizer map:") {
			s2f = ScannerToRanges(scanner)
		}

		if strings.HasPrefix(line, "fertilizer-to-water map:") {
			f2w = ScannerToRanges(scanner)
		}

		if strings.HasPrefix(line, "water-to-light map:") {
			w2l = ScannerToRanges(scanner)
		}

		if strings.HasPrefix(line, "light-to-temperature map:") {
			l2t = ScannerToRanges(scanner)
		}

		if strings.HasPrefix(line, "temperature-to-humidity map:") {
			t2h = ScannerToRanges(scanner)
		}

		if strings.HasPrefix(line, "humidity-to-location map:") {
			h2l = ScannerToRanges(scanner)
		}
    }
	fmt.Println(seeds)
	start := time.Now()
	p2Result := FindLocation(seeds[0])
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println("Started: ", i)
		l := ParallelFind(seeds[i], seeds[i+1])
		if l < p2Result {
			p2Result = l
		}
	}
	elapsed := time.Since(start)
    fmt.Println("Running took: ", elapsed)
	fmt.Println(p2Result)
}

func ParallelFind(start, len int) int {
	result := FindLocation(start)
	for s := start; s < start + len; s++ {
		l := FindLocation(s)
		if l < result {
			result = l
		}
	}
	return result
}
