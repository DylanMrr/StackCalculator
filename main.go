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
		if err == io.EOF {
			break
		}
		text += string(data[:n])

	}

	start := time.Now()
	z := calculate(text)
	//z := calculate("-2+ 1")
	//z := calculate("(1+(4+5+2)-3)+(6+8)")
	//z := calculate("1-(+1+1)")
	fmt.Println(z)

	duration := time.Since(start)
	fmt.Println(duration)
}

func calculate(s string) int {
	output := getExpression(s)
	fmt.Println(output)
	result := count(output)
	return result
}

func count(s string) int {
	stack := Stack{}

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			j := i
			for !isDelimiter(s[i]) && !isOperator(s[i]) {
				i++
				if i == len(s) {
					break
				}
			}
			val, _ := strconv.Atoi(s[j:i])
			stack.push(val)
			i--
		} else if isOperator(s[i]) {
			a, _ := stack.pop().(int)
			if stack.count() == 0 {
				switch string(s[i]) {
				case "+":
					stack.push(a)
				case "-":
					stack.push(-a)
				}
				continue
			}

			b, _ := stack.pop().(int)
			switch string(s[i]) {
			case "+":
				stack.push(b + a)
			case "-":
				stack.push(b - a)
			}
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
			j := i
			for !isOperator(s[i]) && !isDelimiter(s[i]) {
				i++
				if i == len(s) {
					break
				}
			}
			output += s[j:i] + " "
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
