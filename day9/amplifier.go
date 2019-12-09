package main

import (
    // "sync"
    "time"
    "log"
)

type Amplifier struct {
    memory map[int]int
    index  int
    stdout (chan int)
    stdin  (chan int)
    number int
    relative_base int
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
        relative_base: 0,
    }
    return amp
}

func (amp *Amplifier) Compute() {
    for {
        time.Sleep(time.Duration(1) * time.Millisecond)
        log.Println(amp.memory[amp.index], "at index", amp.index)
        param0 := getOpCodeParam(amp.memory[amp.index], 0)
        param1 := getOpCodeParam(amp.memory[amp.index], 1)

        input0 := amp.getValueForParam(param0, 1)
        input1 := amp.getValueForParam(param1, 2)

        opcode := getOpCode(amp.memory[amp.index])
        if opcode == 99 {
            break
        } else if opcode == 1 {
            dest := amp.memory[amp.index + 3]
            param2 := getOpCodeParam(amp.memory[amp.index], 2)
            if param2 == 2 {
                dest = amp.memory[amp.index + 3] + amp.relative_base
            }
            amp.add(input0, input1, dest)
            amp.index += 4
        } else if opcode == 2 {
            dest := amp.memory[amp.index + 3]
            param2 := getOpCodeParam(amp.memory[amp.index], 2)
            if param2 == 2 {
                dest = amp.memory[amp.index + 3] + amp.relative_base
            }
            amp.multiply(input0, input1, dest)
            amp.index += 4
        } else if opcode == 3 {
            dest := amp.memory[amp.index + 1]
            param0 := getOpCodeParam(amp.memory[amp.index], 0)
            if param0 == 2 {
                dest += amp.relative_base
            }
            log.Println(param0, param1)
            log.Println(input0, input1)
            log.Println("relative", amp.relative_base)
            log.Println(dest)
            amp.read(dest)
            amp.index += 2
        } else if opcode == 4 {
            dest := amp.memory[amp.index + 1]
            param0 := getOpCodeParam(amp.memory[amp.index], 0)
            if param0 == 2 {
                dest += amp.relative_base
            }
            amp.print(amp.memory[dest])
            amp.index += 2
        } else if opcode == 5 {
            log.Println("jumping true", input0, input1, param0, param1)
            log.Println("raw", amp.memory[amp.index + 1], amp.memory[amp.index + 2])
            amp.jumpTrue(input0, input1)
        } else if opcode == 6 {
            log.Println("jumping false", input0, input1, param0, param1)
            log.Println("raw", amp.memory[amp.index + 1], amp.memory[amp.index + 2])
            amp.jumpFalse(input0, input1)
        } else if opcode == 7 {
            dest := amp.memory[amp.index + 3]
            param2 := getOpCodeParam(amp.memory[amp.index], 2)
            if param2 == 2 {
                dest = amp.memory[amp.index + 3] + amp.relative_base
            }
            log.Println("performing less than on", input0, input1, "putting result into", dest)
            amp.lessThan(input0, input1, dest)
            amp.index += 4
        } else if opcode == 8 {
            dest := amp.memory[amp.index + 3]
            param2 := getOpCodeParam(amp.memory[amp.index], 2)
            if param2 == 2 {
                dest = amp.memory[amp.index + 3] + amp.relative_base
            }
            log.Println("performing equal to on", input0, input1, "putting result into", dest)
            amp.equalTo(input0, input1, dest)
            amp.index += 4
        } else if opcode == 9 {
            input := amp.memory[amp.index + 1]
            if param0 == 2 {
                input += amp.relative_base
            }
            amp.changeRelative(input)
            amp.index += 2
        } else {
            log.Fatal("HIT UNKNOWN CODE", opcode)
        }
    }
    close(amp.stdout)
    // wg.Done()
}

func (amp *Amplifier) add(value1 int, value2 int, pos int) {
    log.Println("adding", value1, value2, "and putting result into", pos)
    amp.memory[pos] = value1 + value2
}

func (amp *Amplifier) multiply(value1 int, value2 int, pos int) {
    log.Println("multiplying", value1, value2, "and putting result into", pos)
    amp.memory[pos] = value1 * value2
}

func (amp *Amplifier) print(value int) {
    log.Println("output:", value)
    amp.stdout <- value
    // amps[(amp.number + 1) % 5].AddInput(amp.memory[pos])
}

func (amp *Amplifier) read(pos int) {
    amp.memory[pos] = <- amp.stdin
    log.Println("read", amp.memory[pos], "into", pos)
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
    amp.index += 3
    return false
}

func (amp *Amplifier) jumpTrue(value1 int, pos int) bool {
    if value1 != 0 {
        amp.index = pos
        return true
    }
    amp.index += 3
    return false
}

func (amp *Amplifier) changeRelative(value1 int) {
    amp.relative_base += value1
    log.Println("changed relative to", amp.relative_base)
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

func (amp *Amplifier) getValueForParam(param int, paramNum int) int {
    paramIndex := amp.index + paramNum
    if param == 0 {
        if mapHasKey(amp.memory, paramIndex) {
            if mapHasKey(amp.memory, amp.getMemory(paramIndex)) {
                return amp.memory[amp.memory[paramIndex]]
            } else {
                log.Println("got 0 for inner 0")
                return 0
            }
        } else {
            log.Println("got 0 for outer 0")
            return 0
        }
    } else if param == 2 {
        if mapHasKey(amp.memory, paramIndex) {
            return amp.memory[paramIndex] + amp.relative_base
        } else {
            log.Println("got 0 for outer 2")
            return 0
        }
    } else {
        if mapHasKey(amp.memory, paramIndex) {
            return amp.memory[paramIndex]
        }
        return 0
    }
}
