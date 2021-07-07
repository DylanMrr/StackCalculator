package main

type Node struct {
	val  interface{}
	next *Node
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
