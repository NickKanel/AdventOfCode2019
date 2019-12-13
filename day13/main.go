package main

import (
    "strings"
    "strconv"
    "io/ioutil"
    "log"
    "sync"
)

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
    fileBytes, err := ioutil.ReadFile("input")
    check(err)

    codes := parseCodes(string(fileBytes))

    var wg sync.WaitGroup
    wg.Add(1)
    amp := NewIntcodeComputer(codes)
    amp.memory[0] = 2
    go amp.Compute(&wg)

    blocks := 0
    score := 0
    paddleX := 0
    ballX := 0
    for {
        if amp.finished {
            break
        }
        
        value1 := <- amp.stdout
        value2 := <- amp.stdout
        value3 := <- amp.stdout

        if value3 == 2 {
            blocks += 1
        }

        if value3 == 3 {
            paddleX = value1
        }

        if value3 == 4 {
            ballX = value1
            if ballX > paddleX {
                amp.AddInput(1)
            } else if ballX < paddleX {
                amp.AddInput(-1)
            } else {
                amp.AddInput(0)
            }
        }

        if value1 == -1 && value2 == 0 {
            score = value3
        }
    }

    amp.Shutdown()

    log.Println("Blocks:", blocks)
    log.Println("Score:", score)

    wg.Wait()
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
