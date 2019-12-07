package main

import (
    "fmt"
    "time"
    "sync"
    "log"
)

type Amplifier struct {
    memory []int
    index  int
    stdout (chan int)
    stdin  (chan int)
}

func NewAmplifier(codes []int, initialInput int) Amplifier {
    codesCopy := make([]int, len(codes))
    copy(codesCopy, codes)
    amp := Amplifier {
        memory: codes,
        index: 0,
        stdout: make(chan int, 100000),
        stdin: make(chan int, 100000),
    }
    amp.AddInput(initialInput)
    return amp
}

func (amp *Amplifier) Compute(wg *sync.WaitGroup) {
    for {
        log.Println(amp.index)

        param0 := getOpCodeParam(amp.memory[amp.index], 0)
        param1 := getOpCodeParam(amp.memory[amp.index], 1)
        param2 := getOpCodeParam(amp.memory[amp.index], 2)

        opcode := getOpCode(amp.memory[amp.index])
        if opcode == 99 {
            break
        } else if opcode == 1 {
            if param2 == 1 {
                log.Println("THIS SHOULD NEVER HAPPEN FOR ADD")
            }
            input0 := amp.getValueForParam(param0, 1)
            input1 := amp.getValueForParam(param1, 2)
            input2 := amp.getValueForParam(param2, 3)
            amp.add(input0, input1, input2)
            amp.index += 4
            log.Println("incremented")
        } else if opcode == 2 {
            if param2 == 1 {
                log.Println("THIS SHOULD NEVER HAPPEN FOR MULTIPLY")
            }
            input0 := amp.getValueForParam(param0, 1)
            input1 := amp.getValueForParam(param1, 2)
            input2 := amp.getValueForParam(param2, 3)
            amp.multiply(input0, input1, input2)
            amp.index += 4
            log.Println("incremented")
        } else if opcode == 3 {
            input0 := amp.getValueForParam(param0, 1)
            amp.read(input0)
            amp.index += 2
            log.Println("incremented")
        } else if opcode == 4 {
            input0 := amp.getValueForParam(param0, 1)
            amp.print(input0)
            amp.index += 2
            log.Println("incremented")
        } else if opcode == 5 {
            input0 := amp.getValueForParam(param0, 1)
            input1 := amp.getValueForParam(param1, 2)
            if !amp.jumpTrue(input0, input1) {
                amp.index += 3
                log.Println("incremented")
            }
        } else if opcode == 6 {
            input0 := amp.getValueForParam(param0, 1)
            input1 := amp.getValueForParam(param1, 2)
            if !amp.jumpFalse(input0, input1) {
                amp.index += 3
                log.Println("incremented")
            }
        } else if opcode == 7 {
            if param2 == 1 {
                log.Println("THIS SHOULD NEVER HAPPEN FOR ET")
            }
            input0 := amp.getValueForParam(param0, 1)
            input1 := amp.getValueForParam(param1, 2)
            input2 := amp.getValueForParam(param2, 3)
            amp.lessThan(input0, input1, input2)
            amp.index += 4
            log.Println("incremented")
        } else if opcode == 8 {
            if param2 == 1 {
                log.Println("THIS SHOULD NEVER HAPPEN FOR ET")
            }
            input0 := amp.getValueForParam(param0, 1)
            input1 := amp.getValueForParam(param1, 2)
            input2 := amp.getValueForParam(param2, 3)
            amp.equalTo(input0, input1, input2)
            amp.index += 4
            log.Println("incremented")
        } else {
            log.Println("HIT CODE", opcode)
        }
        log.Println("waht")
        time.Sleep(time.Duration(100) * time.Millisecond)
    }
    wg.Done()
}

func (amp *Amplifier) add(value1 int, value2 int, pos int) {
    amp.memory[pos] = value1 + value2
}

func (amp *Amplifier) multiply(value1 int, value2 int, pos int) {
    amp.memory[pos] = value1 * value2
}

func (amp *Amplifier) print(value int) {
    fmt.Println(value)
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
    } else {
        amp.index += 3
    }
    return false
}

func (amp *Amplifier) jumpTrue(value1 int, pos int) bool {
    if value1 != 0 {
        amp.index = pos
        return true
    } else {
        amp.index += 3
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
    if param == 0 {
        return amp.memory[amp.memory[amp.index + paramNum]]
    } else {
        return amp.memory[amp.index + paramNum]
    }
}