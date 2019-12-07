package main

import (
	"strings"
    "strconv"
	"io/ioutil"
    prmt "github.com/gitchander/permutation"
    "sync"
    "log"
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

    perms := makePerms(0, 4)
    for _, perm := range perms {
        log.Println(perm)
        amps := make([]Amplifier, 0)
        for i := 0; i < 5; i += 1 {
            amp := NewAmplifier(codes, perm[i])
            amps = append(amps, amp)
        }

        var wg sync.WaitGroup
        for _, amp := range amps {
            wg.Add(1)
            go amp.Compute(&wg)
        }
        wg.Wait()
        log.Println("done")

        amp5Out := amps[4].ReadOutput()
        log.Println(amp5Out)
    }
}

func makePerms(start int, stop int) [][]int {
    perms := make([][]int, 0)

    stop += 1
    elements := make([]int, stop-start)
    for i := start; i < stop; i++ {
        elements[i] = i
    }

    p := prmt.New(prmt.IntSlice(elements))
    for p.Next() {
        tmp := make([]int, len(elements))
        copy(tmp, elements)
        perms = append(perms, tmp)
    }

    return perms
}