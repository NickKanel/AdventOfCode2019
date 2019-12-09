package main

import (
    "strings"
    "strconv"
    "io/ioutil"
    "sync"
    "log"
)

var amps []*IntcodeComputer

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
    amp := NewIntcodeComputer(codes, 0)
    amp.AddInput(2)
    amp.Compute(&wg)
    close(amp.stdout)
    wg.Wait()
    for value := range amp.stdout {
        log.Println(value)
    }
}
