package main

import (
	//"strings"

	"fmt"
	"strconv"
)

func main() {
	z := calculate("+48 + -48")
	//z := calculate("(1+(4+5+2)-3)+(6+8)")
	//z := calculate("1-(+1+1)")
	fmt.Print(z)
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
			if string(s[i]) == "(" {
				stack.push(string(s[i]))
				if (string(s[i+1])) == "+" {
					i++
				}
			} else if string(s[i]) == ")" {
				char := stack.pop().(string)
				for char != "(" {
					output += char + " "
					char = stack.pop().(string)
				}
			} else {
				if stack.count() > 0 {
					if getPriority(string(s[i])) <= getPriority(stack.peek().(string)) {
						output += stack.pop().(string) + " "
					}
				}
				stack.push(string(s[i]))
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

type Node struct {
	val  interface{}
	next *Node
	//prev *Node
}

type Stack struct {
	top        *Node
	nodesCount int
}

func (this *Stack) stack() Stack {
	node := Node{}
	stack := Stack{top: &node}
	return stack
}

func (this *Stack) push(val interface{}) {
	node := Node{val: val, next: (*this).top}
	(*this).top = &node
	(*this).nodesCount++
}

func (this *Stack) pop() interface{} {
	if (*this).count() == 0 {
		return nil
	}
	temp := (*this).top
	(*this).top = (*this).top.next
	(*this).nodesCount--
	return temp.val
}

func (this *Stack) peek() interface{} {
	if (*this).count() == 0 {
		return nil
	}
	return (*this).top.val
}

func (this *Stack) count() int {
	return (*this).nodesCount
}
