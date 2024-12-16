package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Robot struct {
	warehouse [][]byte
	I, J      int // Size
	i, j      int // Pos

}

func (r *Robot) print() {
	for i := range r.I {
		fmt.Println(string(r.warehouse[i]))
	}
	fmt.Println()
}

func (r *Robot) load(scanner *bufio.Scanner) {
	r.warehouse = make([][]byte, 0)

	for scanner.Scan() {
		// do something with a line
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		r.warehouse = append(r.warehouse, []byte(line))
	}
	r.I = len(r.warehouse)
	r.J = len(r.warehouse[0])

	for i := range r.I {
		for j := range r.J {
			if r.warehouse[i][j] == '@' {
				r.i = i
				r.j = j
				break
			}
		}
	}
}

func (r *Robot) right() {
	canMove := false
	j := r.j + 1
	for ; ; j++ {
		if r.warehouse[r.i][j] == '.' {
			canMove = true
			break
		}
		if r.warehouse[r.i][j] == '#' {
			break
		}
	}
	if canMove {
		for k := j; k > r.j; k-- {
			r.warehouse[r.i][k] = r.warehouse[r.i][k-1]
		}
		r.warehouse[r.i][r.j] = '.'
		r.j += 1
	}
}

func (r *Robot) left() {
	canMove := false
	j := r.j - 1
	for ; ; j-- {
		if r.warehouse[r.i][j] == '.' {
			canMove = true
			break
		}
		if r.warehouse[r.i][j] == '#' {
			break
		}
	}
	if canMove {
		for k := j; k < r.j; k++ {
			r.warehouse[r.i][k] = r.warehouse[r.i][k+1]
		}
		r.warehouse[r.i][r.j] = '.'
		r.j -= 1
	}
}

func (r *Robot) up() {
	canMove := false
	i := r.i - 1

	for ; ; i-- {
		if r.warehouse[i][r.j] == '.' {
			canMove = true
			break
		}
		if r.warehouse[i][r.j] == '#' {
			break
		}
	}
	if canMove {
		for k := i; k < r.i; k++ {
			r.warehouse[k][r.j] = r.warehouse[k+1][r.j]
		}
		r.warehouse[r.i][r.j] = '.'
		r.i -= 1
	}
}

func (r *Robot) down() {
	canMove := false
	i := r.i + 1

	for ; ; i++ {
		if r.warehouse[i][r.j] == '.' {
			canMove = true
			break
		}
		if r.warehouse[i][r.j] == '#' {
			break
		}
	}
	if canMove {
		for k := i; k > r.i; k-- {
			r.warehouse[k][r.j] = r.warehouse[k-1][r.j]
		}
		r.warehouse[r.i][r.j] = '.'
		r.i += 1
	}
}

func (r *Robot) sumBoxes() int {
	sum := 0
	for i := range r.I {
		for j := range r.J {
			if r.warehouse[i][j] == 'O' {
				sum += 100*i + j
			}
		}
	}
	return sum
}

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

	var r Robot
	r.load(scanner)
	r.print()

	for scanner.Scan() {
		// do something with a line
		line := []byte(scanner.Text())
		for _, b := range line {
			if b == '<' {
				r.left()
			} else if b == '>' {
				r.right()
			} else if b == '^' {
				r.up()
			} else if b == 'v' {
				r.down()
			} else {
				fmt.Println("Unknown command:", b)
			}
		}
	}
	r.print()
	fmt.Println(r.sumBoxes())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
