package main

import (
    "sync"
    "log"
)

type Amplifier struct {
    memory []int
    index  int
    stdout (chan int)
    stdin  (chan int)
    number int
}

func NewAmplifier(codes []int, number int, initialInput int) Amplifier {
    codesCopy := make([]int, len(codes))
    copy(codesCopy, codes)
    amp := Amplifier {
        memory: codesCopy,
        index: 0,
        stdout: make(chan int, 100000),
        stdin: make(chan int, 100000),
        number: number,
    }
    amp.AddInput(initialInput)
    return amp
}

func (amp *Amplifier) Compute(wg *sync.WaitGroup) {
    for {
        param0 := getOpCodeParam(amp.memory[amp.index], 0)
        param1 := getOpCodeParam(amp.memory[amp.index], 1)

        input0 := amp.getValueForParam(param0, 1)
        input1 := amp.getValueForParam(param1, 2)

        opcode := getOpCode(amp.memory[amp.index])
        if opcode == 99 {
            break
        } else if opcode == 1 {
            dest := amp.memory[amp.index + 3]
            amp.add(input0, input1, dest)
            amp.index += 4
        } else if opcode == 2 {
            dest := amp.memory[amp.index + 3]
            amp.multiply(input0, input1, dest)
            amp.index += 4
        } else if opcode == 3 {
            dest := amp.memory[amp.index + 1]
            amp.read(dest)
            amp.index += 2
        } else if opcode == 4 {
            addr := amp.memory[amp.index + 1]
            amp.print(addr)
            amp.index += 2
        } else if opcode == 5 {
            amp.jumpTrue(input0, input1)
        } else if opcode == 6 {
            amp.jumpFalse(input0, input1)
        } else if opcode == 7 {
            dest := amp.memory[amp.index + 3]
            amp.lessThan(input0, input1, dest)
            amp.index += 4
        } else if opcode == 8 {
            dest := amp.memory[amp.index + 3]
            amp.equalTo(input0, input1, dest)
            amp.index += 4
        } else {
            log.Fatal("HIT UNKNOWN CODE", opcode)
        }
    }
    wg.Done()
}

func (amp *Amplifier) add(value1 int, value2 int, pos int) {
    amp.memory[pos] = value1 + value2
}

func (amp *Amplifier) multiply(value1 int, value2 int, pos int) {
    amp.memory[pos] = value1 * value2
}

func (amp *Amplifier) print(pos int) {
    amps[(amp.number + 1) % 5].AddInput(amp.memory[pos])
}

func (amp *Amplifier) read(pos int) {
    amp.memory[pos] = <- amp.stdin
}

func (amp *Amplifier) lessThan(value1 int, value2 int, pos int) {
    if value1 < value2 {
        amp.memory[pos] = 1
    } else {
        amp.memory[pos] = 0
    }
}

func (amp *Amplifier) equalTo(value1 int, value2 int, pos int) {
    if value1 == value2 {
        amp.memory[pos] = 1
    } else {
        amp.memory[pos] = 0
    }
}

func (amp *Amplifier) jumpFalse(value1 int, pos int) bool {
    if value1 == 0 {
        amp.index = pos
        return true
    }
    return false
}

func (amp *Amplifier) jumpTrue(value1 int, pos int) bool {
    if value1 != 0 {
        amp.index = pos
        return true
    }
    return false
}

func (amp *Amplifier) AddInput(value int) {
    amp.stdin <- value
}

func (amp *Amplifier) ReadOutput() int {
    return <- amp.stdout
}

func (amp *Amplifier) getValueForParam(param int, paramNum int) int {
    paramIndex := amp.index + paramNum
    if param == 0 {
        if paramIndex < len(amp.memory) {
            if amp.memory[amp.index + paramNum] < len(amp.memory) {
                return amp.memory[amp.memory[paramIndex]]
            }
        }
        return -1
    } else {
        return amp.memory[paramIndex]
    }
}
