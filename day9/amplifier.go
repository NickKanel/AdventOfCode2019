package main

import (
    "sync"
    // "time"
    "log"
)

type Amplifier struct {
    memory map[int]int
    index  int
    stdout (chan int)
    stdin  (chan int)
    number int
    relativeBase int
}

func NewAmplifier(codes []int, number int) Amplifier {
    codesCopy := make([]int, len(codes))
    copy(codesCopy, codes)

    memory := make(map[int]int)
    for i, j := range codesCopy {
        memory[i] = j
    }

    amp := Amplifier {
        memory: memory,
        index: 0,
        stdout: make(chan int, 100000),
        stdin: make(chan int, 100000),
        number: number,
        relativeBase: 0,
    }
    return amp
}

func (amp *Amplifier) Compute(wg *sync.WaitGroup) {
    for {
        param1 := getOpCodeParam(amp.memory[amp.index], 0)
        param2 := getOpCodeParam(amp.memory[amp.index], 1)
        param3 := getOpCodeParam(amp.memory[amp.index], 2)

        addr1 := amp.getAddressForParam(param1, 1)
        addr2 := amp.getAddressForParam(param2, 2)
        addr3 := amp.getAddressForParam(param3, 3)

        opcode := getOpCode(amp.memory[amp.index])
        if opcode == 99 {
            break
        } else if opcode == 1 {
            amp.add(addr1, addr2, addr3)
            amp.index += 4
        } else if opcode == 2 {
            amp.multiply(addr1, addr2, addr3)
            amp.index += 4
        } else if opcode == 3 {
            amp.read(addr1)
            amp.index += 2
        } else if opcode == 4 {
            amp.print(addr1)
            amp.index += 2
        } else if opcode == 5 {
            amp.jumpTrue(addr1, addr2)
        } else if opcode == 6 {
            amp.jumpFalse(addr1, addr2)
        } else if opcode == 7 {
            amp.lessThan(addr1, addr2, addr3)
            amp.index += 4
        } else if opcode == 8 {
            amp.equalTo(addr1, addr2, addr3)
            amp.index += 4
        } else if opcode == 9 {
            amp.changeRelative(addr1)
            amp.index += 2
        } else {
            log.Fatal("HIT UNKNOWN CODE", opcode)
        }
    }
    close(amp.stdout)
    wg.Done()
}

func (amp *Amplifier) add(addr1 int, addr2 int, addr3 int) {
    // log.Println("adding", amp.getMemory(addr1), amp.getMemory(addr2), "and putting result into", addr3)
    amp.memory[addr3] = amp.getMemory(addr1) + amp.getMemory(addr2)
}

func (amp *Amplifier) multiply(addr1 int, addr2 int, addr3 int) {
    // log.Println("multiplying", amp.getMemory(addr1), amp.getMemory(addr2), "and putting result into", addr3)
    amp.memory[addr3] = amp.getMemory(addr1) * amp.getMemory(addr2)
}

func (amp *Amplifier) print(addr1 int) {
    // log.Println("output:", amp.getMemory(addr1))
    amp.stdout <- amp.getMemory(addr1)
    // amps[(amp.number + 1) % 5].AddInput(amp.memory[pos])
}

func (amp *Amplifier) read(addr1 int) {
    amp.memory[addr1] = <- amp.stdin
    // log.Println("read", amp.getMemory(addr1), "into", addr1)
}

func (amp *Amplifier) jumpTrue(addr1 int, addr2 int) {
    if amp.getMemory(addr1) != 0 {
        // log.Println("jumping to", amp.getMemory(addr2))
        amp.index = amp.getMemory(addr2)
    } else {
        amp.index += 3
        // log.Println("not jumping, incremented to", amp.index)
    }
}

func (amp *Amplifier) jumpFalse(addr1 int, addr2 int) {
    if amp.getMemory(addr1) == 0 {
        // log.Println("jumping to", amp.getMemory(addr2))
        amp.index = amp.getMemory(addr2)
    } else {
        amp.index += 3
        // log.Println("not jumping, incremented to", amp.index)
    }
}

func (amp *Amplifier) lessThan(addr1 int, addr2 int, addr3 int) {
    if amp.getMemory(addr1) < amp.getMemory(addr2) {
        amp.memory[addr3] = 1
    } else {
        amp.memory[addr3] = 0
    }
}

func (amp *Amplifier) equalTo(addr1 int, addr2 int, addr3 int) {
    if amp.getMemory(addr1) == amp.getMemory(addr2) {
        amp.memory[addr3] = 1
    } else {
        amp.memory[addr3] = 0
    }
}

func (amp *Amplifier) changeRelative(addr1 int) {
    amp.relativeBase += amp.getMemory(addr1)
    // log.Println("changed relative to", amp.relativeBase)
}

func (amp *Amplifier) AddInput(value int) {
    amp.stdin <- value
}

func (amp *Amplifier) ReadOutput() int {
    return <- amp.stdout
}

func (amp *Amplifier) getMemory(i int) int {
    if mapHasKey(amp.memory, i) {
        return amp.memory[i]
    }
    return 0
}

func (amp *Amplifier) getAddressForParam(param int, paramNum int) int {
    paramIndex := amp.index + paramNum
    if param == 0 {
        return amp.getMemory(paramIndex)
    } else if param == 2 {
        return amp.getMemory(paramIndex) + amp.relativeBase
    } else {
        return paramIndex
    }
}
