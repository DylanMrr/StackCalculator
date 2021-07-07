package main

import (
	"strconv"
	"unicode/utf8"
)

var prior map[string]int = map[string]int{
	"(": 0,
	")": 1,
	"+": 2,
	"-": 3,
}

var bufPlus []byte
var bufMinus []byte
var bufOpen []byte
var bufClose []byte

func configure() {
	bufPlus = make([]byte, 1)
	bufMinus = make([]byte, 1)
	bufOpen = make([]byte, 1)
	bufClose = make([]byte, 1)
	_ = utf8.EncodeRune(bufPlus, '+')
	_ = utf8.EncodeRune(bufMinus, '-')
	_ = utf8.EncodeRune(bufOpen, '(')
	_ = utf8.EncodeRune(bufClose, ')')
}

func getPriority(s string) int {
	val, ok := prior[s]
	if ok {
		return val
	}
	return 4
}

func isDelimiter(s byte) bool {
	return string(s) == " "
}

func isDigit(s byte) bool {
	if _, err := strconv.Atoi(string(s)); err == nil {
		return true
	}
	return false
}

func isOperator(s byte) bool {
	operators := []byte{bufPlus[0], bufMinus[0], bufOpen[0], bufClose[0]}
	for _, value := range operators {
		if value == s {
			return true
		}
	}
	return false
}
