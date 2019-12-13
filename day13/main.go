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
    amp.AddInput(1)
    go amp.Compute(&wg)

    blocks := 0
    for {
        if amp.finished {
            break
        }
        
        _ = <- amp.stdout
        _ = <- amp.stdout
        value3 := <- amp.stdout

        if value3 == 2 {
            blocks += 1
        }
    }

    amp.Shutdown()

    log.Println(blocks)

    wg.Wait()
}
