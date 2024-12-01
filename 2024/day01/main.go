package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MinHeap struct {
	values []int
	size   int
}

func (mh *MinHeap) swap(i1, i2 int) {
	if i1 >= mh.size || i2 >= mh.size {
		errors.New("Wrong index")
	}
	value := mh.values[i1]
	mh.values[i1] = mh.values[i2]
	mh.values[i2] = value
}

func (mh *MinHeap) push(v int) {
	mh.values = append(mh.values, v)
	mh.size += 1
	for index := mh.size - 1; index > 0; {
		parent := index / 2
		if mh.values[parent] > mh.values[index] {
			mh.swap(index, parent)
			index = parent
		} else {
			break
		}
	}

}

func (mh *MinHeap) pop() int {
	value := mh.values[0]
	mh.values[0] = mh.values[mh.size-1]
	mh.size -= 1
	for index := 0; ; {
		child1 := index * 2
		child2 := child1 + 1

		if child1 >= mh.size {
			break
		}
		minIndex := child1
		if child2 < mh.size && mh.values[child2] < mh.values[minIndex] {
			minIndex = child2
		}
		if mh.values[minIndex] < mh.values[index] {
			mh.swap(index, minIndex)
			index = minIndex
		} else {
			break
		}
	}
	return value
}

func (mh *MinHeap) isEmpty() bool {
	return mh.size == 0
}

func main() {
	var left MinHeap
	var right MinHeap
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
		v1, _ := strconv.Atoi(row[0])
		left.push(v1)
		v2, _ := strconv.Atoi(row[1])
		right.push(v2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for !left.isEmpty() && !right.isEmpty() {
		v1 := left.pop()
		v2 := right.pop()
		diff := v1 - v2
		if diff >= 0 {
			sum += diff
		} else {
			sum -= diff
		}
	}
	fmt.Println(sum)
}
