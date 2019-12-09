package main

import (
    "strings"
    "strconv"
    "io/ioutil"
    prmt "github.com/gitchander/permutation"
    "sync"
    "log"
)

var amps []*Amplifier

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

    // maxOut := 0
    // perms := makePerms(5, 9)


    var wg sync.WaitGroup
    wg.Add(1)
    amp := NewAmplifier(codes, 0)
    amp.AddInput(2)
    amp.Compute(&wg)
    wg.Wait()
    for value := range amp.stdout {
        log.Println(value)
    }
    // output := <- amp.stdout



    // for _, perm := range perms {
    //     amps = make([]*Amplifier, 0)
    //     for i := 0; i < 5; i += 1 {
    //         amp := NewAmplifier(codes, i, perm[i])
    //         amps = append(amps, &amp)
    //     }

    //     amps[0].AddInput(0)

    //     var wg sync.WaitGroup
    //     for _, amp := range amps {
    //         wg.Add(1)
    //         go amp.Compute(&wg)
    //     }
    //     wg.Wait()

    //     amp5Out := <- amps[0].stdin

    //     if amp5Out > maxOut {
    //         maxOut = amp5Out
    //     }
    // }

    // log.Println(output)
}

func makePerms(start int, stop int) [][]int {
    var perms [][]int

    stop += 1
    var elements []int
    for i := start; i < stop; i++ {
        elements = append(elements, i)
    }

    p := prmt.New(prmt.IntSlice(elements))
    for p.Next() {
        tmp := make([]int, len(elements))
        copy(tmp, elements)
        perms = append(perms, tmp)
    }

    return perms
}