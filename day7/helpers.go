package main

import (
	"strconv"
)

func leftPad(s string, l int, c rune) string {
	if len(s) >= l {
		return s
	}

	sLen := len(s)
	for i := 0; i < (l - sLen); i++ {
		s = string(c) + s
	}

	return s
}

func getOpCode(code int) int {
	paddedCode := leftPad(strconv.Itoa(code), 2, rune(48))
	intVal, err := strconv.Atoi(paddedCode[len(paddedCode) - 2:])
	check(err)
	return intVal
}

func getOpCodeParam(code int, index int) int {
	paddedCode := leftPad(strconv.Itoa(code), 2 + index + 1, rune(48))
	params := paddedCode[:len(paddedCode) - 2]
	intVal, err := strconv.Atoi(string([]rune(paddedCode)[len(params) - 1 - index]))
	check(err)
	return intVal
}
