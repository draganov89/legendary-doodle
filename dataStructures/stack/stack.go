package stack

import (
	"fmt"
	"strings"
)

type Node struct { 
	Value interface{}
	Next *Node
}

type Stack struct {
	top *Node
	count int
}

// Stack constructor
func ConstructStack() *Stack {
	return &Stack{nil, 0}
}

func (t Stack) Count() int {
	return t.count
}
func (t *Stack) Empty() {
	t.count = 0
	t.top = nil
}

func (t *Stack) Push(val interface{}) {
	node := &Node{val, nil}

	node.Next = t.top
	t.top = node

	t.count++
}

func (t *Stack) Pull() (interface{}, bool) {
	if t.top == nil {
		return nil, false
	}
	result := t.top.Value
	t.top = t.top.Next
	t.count--
	return result, true
}

func (t *Stack) Pop() (interface{}, bool) {
	if t.top == nil {
		return nil, false
	}
	return t.top.Value, true
}

func (t Stack) String() string {
	current := t.top
	var result string
	for current != nil {
		result += fmt.Sprint(current.Value)
		result += " "
		current = current.Next
	}
	return strings.Trim(result, " ")
}