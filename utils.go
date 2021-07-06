package main

import "strconv"

var prior map[string]int = map[string]int{
	"(": 0,
	")": 1,
	"+": 2,
	"-": 3,
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
	operators := []string{"+", "-", "(", ")"}
	for _, value := range operators {
		if value == string(s) {
			return true
		}
	}
	return false
}
