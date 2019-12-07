package main

import (
	"strconv"
	// "log"
)

func leftPad(s string, l int, c rune) string {
	if len(s) >= l {
		return s
	}

	sLen := len(s)
	sRunes := []rune(s)
	for i := 0; i < (l - sLen); i++ {
		sRunes = append([]rune{c}, sRunes...)
	}

	return string(sRunes)
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

// def left_pad(string, length, pad = '0'):
//     if len(string) >= length:
//         return string
//     return ''.join([pad]*(length-len(string))) + string

// def get_op_code(code):
//     return int(left_pad(str(code), 2)[-2:])

// def get_op_code_param(code, index):
//     padded_code = left_pad(str(code), 2 + index + 1)
//     params = padded_code[:len(padded_code)-2]
//     return int(params[len(params) - 1 - index])