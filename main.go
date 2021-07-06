package main

import (
	//"strings"

	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	var text string
	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		text += string(data[:n])
	}

	z := calculate(text)
	//z := calculate("2147483647")
	//z := calculate("(1+(4+5+2)-3)+(6+8)")
	//z := calculate("1-(+1+1)")
	fmt.Println(z)

	duration := time.Since(start)
	fmt.Println(duration)
}

func calculate(s string) int {
	/*index := 0
	  for strings.HasPrefix(s, " "){
	      index++
	  }*/
	output := getExpression(s)
	fmt.Println(output)
	result := count(output)
	return result
}

func count(s string) int {
	var result int
	stack := Stack{}

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			var temp string

			for !isDelimiter(s[i]) && !isOperator(s[i]) {
				temp += string(s[i])
				i++
				if i == len(s) {
					break
				}
			}
			val, _ := strconv.Atoi(temp)
			stack.push(int(val))
			i--
		} else if isOperator(s[i]) {
			a, err := stack.pop().(int)
			b, err := stack.pop().(int)

			if !err {
				switch string(s[i]) {
				case "+":
					result = a
				case "-":
					result = -a
				}
				stack.push(result)
				continue
			}

			switch string(s[i]) {
			case "+":
				result = b + a
			case "-":
				result = b - a
			}
			stack.push(result)
		}
	}
	return stack.peek().(int)
}

func getExpression(s string) string {
	var output string
	stack := Stack{}

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if isDelimiter(s[i]) {
			continue
		}
		if isDigit(s[i]) {
			for !isOperator(s[i]) && !isDelimiter(s[i]) {
				output += string(s[i])
				i++
				if i == len(s) {
					break
				}
			}
			output += " "
			i--
		}
		if isOperator(s[i]) {
			if char == "(" {
				stack.push(char)
				if (string(s[i+1])) == "+" {
					i++
				}
			} else if char == ")" {
				char := stack.pop().(string)
				for char != "(" {
					output += char + " "
					char = stack.pop().(string)
				}
			} else {
				if stack.count() > 0 {
					if getPriority(char) <= getPriority(stack.peek().(string)) {
						output += stack.pop().(string) + " "
					}
				}
				stack.push(char)
			}
		}
	}
	for stack.count() > 0 {
		output += stack.pop().(string) + " "
	}
	return output
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

func getPriority(s string) int {
	switch s {
	case "(":
		return 0
	case ")":
		return 1
	case "+":
		return 2
	case "-":
		return 3
	default:
		return 4
	}
}

func isDelimiter(s byte) bool {
	if string(s) == " " {
		return true
	}
	return false
}

func isDigit(s byte) bool {
	if _, err := strconv.Atoi(string(s)); err == nil {
		return true
	}
	return false
}
