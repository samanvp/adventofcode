package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

type Element struct {
    i, j int
    distance int
    fromNorth, fromSouth, fromWest, fromEast bool
}

var elementsToProcess []Element

var pipesMaxI int
var pipesMaxJ int
var pipes [][]rune

var visited map[string]int
func main() {
    // open file
    f, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

   // read the file line by line using scanner 
   scanner := bufio.NewScanner(f)

   startI := -1
   startJ := -1
   pipesMaxI = -1
   pipesMaxJ = -1
   lineI := 0
    for scanner.Scan() {
        // do something with a line
        line := scanner.Text()
        var row []rune
        for j, c := range line {
            row = append(row, c)
            if c == 'S' {
                startI = lineI
                startJ = j
            }
        }
        if pipesMaxI == -1 {
            pipesMaxI = len(row)
        }
        pipes = append(pipes, row)
        lineI += 1 
    }
    pipesMaxJ = len(pipes)
    fmt.Println(pipesMaxI, "  by  ", pipesMaxJ)
    AddNeighbors(startI, startJ)

    visited = make(map[string]int)
    var distance int
    for len(elementsToProcess) > 0 {
        e := elementsToProcess[0]
        elementsToProcess = elementsToProcess[1:]

        if e.i == startI && e.j == startJ {
            fmt.Println("Arrived back at `S`")
            distance = e.distance
            break
        }
        if e.i < 0 || e.j < 0 || e.i >= pipesMaxI || e.j >= pipesMaxJ {
            fmt.Println("skip out of border element: ", e)
            continue
        }
        _, ok := visited[CordToString(e.i, e.j)]
        if !ok {
            AddNext(e)
        } else {
            fmt.Println("skip a node with more one path: ", e)
        }
    }

    fmt.Println(visited)

    p1Result := distance / 2
    if distance % 2 != 0 {
        p1Result += 1
    } 
    fmt.Println(p1Result)
}

func CordToString(i, j int) string {
    return strconv.Itoa(i) + ":" + strconv.Itoa(j)
}

func MoveNorth(e Element) Element {
    return Element{i : e.i - 1, j : e.j, distance : e.distance + 1, fromSouth: true}
}

func MoveSouth(e Element) Element {
    return Element{i : e.i + 1, j : e.j, distance : e.distance + 1, fromNorth: true}
}

func MoveEast(e Element) Element {
    return Element{i : e.i, j : e.j + 1, distance : e.distance + 1, fromWest: true}
}

func MoveWest(e Element) Element {
    return Element{i : e.i, j : e.j - 1, distance : e.distance + 1, fromEast: true}
}

func AddNext(e Element) {
    visited[CordToString(e.i, e.j)] = e.distance

    if pipes[e.i][e.j] == '|' {
        if e.fromNorth {
            elementsToProcess = append(elementsToProcess, MoveSouth(e))
        } else if e.fromSouth {
            elementsToProcess = append(elementsToProcess, MoveNorth(e))
        } else {
            fmt.Println("Wrong entrance :", e)
        }
    }
    if pipes[e.i][e.j] == '-' {
        if e.fromWest {
            elementsToProcess = append(elementsToProcess, MoveEast(e))
        } else if e.fromEast {
            elementsToProcess = append(elementsToProcess, MoveWest(e))
        } else {
            fmt.Println("Wrong entrance :", e)
        }
    }
    // L is a 90-degree bend connecting north and east.
    if pipes[e.i][e.j] == 'L' {
        if e.fromNorth {
            elementsToProcess = append(elementsToProcess, MoveEast(e))
        } else if e.fromEast {
            elementsToProcess = append(elementsToProcess, MoveNorth(e))
        } else {
            fmt.Println("Wrong entrance :", e)
        }
    }
    // J is a 90-degree bend connecting north and west.
    if pipes[e.i][e.j] == 'J' {
        if e.fromNorth {
            elementsToProcess = append(elementsToProcess, MoveWest(e))
        } else if e.fromWest {
            elementsToProcess = append(elementsToProcess, MoveNorth(e))
        } else {
            fmt.Println("Wrong entrance :", e)
        }
    }
    // 7 is a 90-degree bend connecting south and west.
    if pipes[e.i][e.j] == '7' {
        if e.fromWest {
            elementsToProcess = append(elementsToProcess, MoveSouth(e))
        } else if e.fromSouth{
            elementsToProcess = append(elementsToProcess, MoveWest(e))
        } else {
            fmt.Println("Wrong entrance :", e)
        }
    }
    // F is a 90-degree bend connecting south and east.
    if pipes[e.i][e.j] == 'F' {
        if e.fromEast {
            elementsToProcess = append(elementsToProcess, MoveSouth(e))
        } else if e.fromSouth{
            elementsToProcess = append(elementsToProcess, MoveEast(e))
        } else {
            fmt.Println("Wrong entrance :", e)
        }
    }
}

func AddNeighbors(startI, startJ int) {
    if startI > 0 {
        next := pipes[startI - 1][startJ]
        if next == '|' || next == '7' || next == 'F' {
            e := Element{i : startI - 1, j : startJ, distance : 1, fromSouth : true}
            elementsToProcess = append(elementsToProcess, e)
            return
        }    
    }

    if startI < pipesMaxI {
        next := pipes[startI + 1][startJ]
        if next == '|' || next == 'L' || next == 'J' {
            e := Element{i : startI + 1, j : startJ, distance : 1, fromNorth : true}
            elementsToProcess = append(elementsToProcess, e)
            return
        }    
    }

    if startJ > 0 {
        next := pipes[startI][startJ - 1]
        if next == '-' || next == 'F' || next == 'L' {
            e := Element{i : startI, j : startJ - 1, distance : 1, fromWest : true}
            elementsToProcess = append(elementsToProcess, e)
            return
        }    
    }

    if startJ < pipesMaxJ {
        next := pipes[startI][startJ + 1]
        if next == '-' || next == 'J' || next == '7' {
            e := Element{i : startI, j : startJ + 1, distance : 1, fromEast : true}
            elementsToProcess = append(elementsToProcess, e)
            return
        }    
    }
}