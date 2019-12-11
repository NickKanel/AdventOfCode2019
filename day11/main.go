package main

import (
    "strings"
    "strconv"
    "io/ioutil"
    "sync"
    "log"
    // "fmt"
)

var panels [][]int
var robotPos []int
var robotDirection int

func initValues() {
    for i := 0; i < 100; i++ {
        inner := make([]int, 100)
        panels = append(panels, inner)
    }
    robotDirection = 0
    robotPos = []int{49,49}
}

type Point struct {
    x int
    y int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parseCodes(s string) []int {
    intCodes := make([]int, 0)
    splitString := strings.Split(s, ",")
    for _, string := range splitString {
        intValue, err := strconv.Atoi(string)
        check(err)
        intCodes = append(intCodes, intValue)
    }
    return intCodes
}

func main() {
    colored1 := make(map[Point]bool)
    fileBytes, err := ioutil.ReadFile("input")
    check(err)

    codes := parseCodes(string(fileBytes))

    initValues()

    var wg sync.WaitGroup
    wg.Add(1)
    amp := NewIntcodeComputer(codes)
    amp.AddInput(0)
    go amp.Compute(&wg)
    first := true
    for value := range amp.stdout {
        if first {
            if value == 1 {
                colored1[Point{robotPos[0],robotPos[1]}] = true
            }
            setColor(robotPos, value)
            if amp.finished {
                log.Println("Number of points colored 1 is", len(colored1))
                // for _, row := range panels {
                //     for _, col := range row {
                //         fmt.Print(col)
                //     }
                //     fmt.Println()
                // }
            }
        } else {
            if !amp.finished {
                turnRobot(value)
                moveRobot()
                amp.AddInput(getColor(robotPos))
            }
        }
        first = !first
    }
    wg.Wait()
}

func turnRobot(value int) {
    if value == 0 {
        if robotDirection == 0 {
            robotDirection = 3
        } else {
            robotDirection -= 1
        }
    } else {
        if robotDirection == 3 {
            robotDirection = 0
        } else {
            robotDirection += 1
        }
    }
}

func moveRobot() {
    if robotDirection == 0 {
        robotPos = []int{robotPos[0], robotPos[1] - 1}
    } else if robotDirection == 1 {
        robotPos = []int{robotPos[0] + 1, robotPos[1]}
    } else if robotDirection == 2 {
        robotPos = []int{robotPos[0], robotPos[1] + 1}
    } else {
        robotPos = []int{robotPos[0] - 1, robotPos[1]}
    }
}

func getColor(pos []int) int {
    return panels[pos[0]][pos[1]]
}

func setColor(pos []int, color int) {
    panels[pos[0]][pos[1]] = color
}