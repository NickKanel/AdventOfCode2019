package main

import (
    "sync"
    // "time"
    "log"
)

type IntcodeComputer struct {
    memory map[int]int
    index  int
    stdout (chan int)
    stdin  (chan int)
    relativeBase int
    finished bool
}

func NewIntcodeComputer(codes []int) IntcodeComputer {
    memory := make(map[int]int)
    for i, j := range codes {
        memory[i] = j
    }

    comp := IntcodeComputer {
        memory: memory,
        index: 0,
        relativeBase: 0,
        stdout: make(chan int, 100000),
        stdin: make(chan int, 100000),
        finished: false,
    }
    return comp
}

func (comp *IntcodeComputer) Compute(wg *sync.WaitGroup) {
    for {
        param1 := getOpCodeParam(comp.memory[comp.index], 0)
        param2 := getOpCodeParam(comp.memory[comp.index], 1)
        param3 := getOpCodeParam(comp.memory[comp.index], 2)

        addr1 := comp.getAddressForParam(param1, 1)
        addr2 := comp.getAddressForParam(param2, 2)
        addr3 := comp.getAddressForParam(param3, 3)

        opcode := getOpCode(comp.memory[comp.index])
        if opcode == 99 {
            break
        } else if opcode == 1 {
            comp.add(addr1, addr2, addr3)
            comp.index += 4
        } else if opcode == 2 {
            comp.multiply(addr1, addr2, addr3)
            comp.index += 4
        } else if opcode == 3 {
            comp.read(addr1)
            comp.index += 2
        } else if opcode == 4 {
            comp.print(addr1)
            comp.index += 2
        } else if opcode == 5 {
            comp.jumpTrue(addr1, addr2)
        } else if opcode == 6 {
            comp.jumpFalse(addr1, addr2)
        } else if opcode == 7 { 
            comp.lessThan(addr1, addr2, addr3)
            comp.index += 4
        } else if opcode == 8 {
            comp.equalTo(addr1, addr2, addr3)
            comp.index += 4
        } else if opcode == 9 {
            comp.changeRelative(addr1)
            comp.index += 2
        } else {
            log.Fatal("HIT UNKNOWN CODE ", opcode)
        }
    }
    comp.finished = true
    close(comp.stdin)
    close(comp.stdout)
    wg.Done()
}

func (comp *IntcodeComputer) add(addr1 int, addr2 int, addr3 int) {
    // log.Println("adding", comp.getMemory(addr1), comp.getMemory(addr2), "and putting result into", addr3)
    comp.memory[addr3] = comp.getMemory(addr1) + comp.getMemory(addr2)
}

func (comp *IntcodeComputer) multiply(addr1 int, addr2 int, addr3 int) {
    // log.Println("multiplying", comp.getMemory(addr1), comp.getMemory(addr2), "and putting result into", addr3)
    comp.memory[addr3] = comp.getMemory(addr1) * comp.getMemory(addr2)
}

func (comp *IntcodeComputer) print(addr1 int) {
    comp.stdout <- comp.getMemory(addr1)
    // comps[(comp.number + 1) % 5].AddInput(comp.memory[pos])
}

func (comp *IntcodeComputer) read(addr1 int) {
    comp.memory[addr1] = <- comp.stdin
    // log.Println("read", comp.getMemory(addr1), "into", addr1)
}

func (comp *IntcodeComputer) jumpTrue(addr1 int, addr2 int) {
    if comp.getMemory(addr1) != 0 {
        // log.Println("jumping to", comp.getMemory(addr2))
        comp.index = comp.getMemory(addr2)
    } else {
        comp.index += 3
        // log.Println("not jumping, incremented to", comp.index)
    }
}

func (comp *IntcodeComputer) jumpFalse(addr1 int, addr2 int) {
    if comp.getMemory(addr1) == 0 {
        // log.Println("jumping to", comp.getMemory(addr2))
        comp.index = comp.getMemory(addr2)
    } else {
        comp.index += 3
        // log.Println("not jumping, incremented to", comp.index)
    }
}

func (comp *IntcodeComputer) lessThan(addr1 int, addr2 int, addr3 int) {
    if comp.getMemory(addr1) < comp.getMemory(addr2) {
        comp.memory[addr3] = 1
    } else {
        comp.memory[addr3] = 0
    }
}

func (comp *IntcodeComputer) equalTo(addr1 int, addr2 int, addr3 int) {
    if comp.getMemory(addr1) == comp.getMemory(addr2) {
        comp.memory[addr3] = 1
    } else {
        comp.memory[addr3] = 0
    }
}

func (comp *IntcodeComputer) changeRelative(addr1 int) {
    comp.relativeBase += comp.getMemory(addr1)
    // log.Println("changed relative to", comp.relativeBase)
}

func (comp *IntcodeComputer) getMemory(i int) int {
    if mapHasKey(comp.memory, i) {
        return comp.memory[i]
    }
    return 0
}

func (comp *IntcodeComputer) AddInput(value int) {
    comp.stdin <- value
}

func (comp *IntcodeComputer) ReadOutput() int {
    return <- comp.stdout
}

func (comp *IntcodeComputer) getAddressForParam(param int, paramNum int) int {
    paramIndex := comp.index + paramNum
    if param == 0 {
        return comp.getMemory(paramIndex)
    } else if param == 2 {
        return comp.getMemory(paramIndex) + comp.relativeBase
    } else {
        return paramIndex
    }
}
