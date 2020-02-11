package queue

import (
	"fmt"
	"strings"
)

type node struct { 
	Value interface{}
	Next *node
}

type Queue struct {
	head, tail *node
	count int
}

func (q Queue) Count() int {
	return q.count
}

func ConstructQueue() *Queue {
	return &Queue{nil, nil, 0}
}

func (t *Queue) Empty() {
	t.head = nil
	t.tail = nil
	t.count = 0
}

func (t *Queue) Enqueue(val interface{}) {
	node := &node{Value: val, Next: nil}

	if t.tail == nil {
		t.head = node
		t.tail = node
	} else {
		t.tail.Next = node
		t.tail = node
	}
	t.count++
}

func (t *Queue) Dequeue() (interface{}, bool){
	if t.head == nil {
		return nil, false
	} 
	result := t.head.Value
	if t.head == t.tail {
		t.tail, t.head = nil, nil
	} else {
		t.head = t.head.Next
	}
	t.count--
	return result, true
}

func (t Queue) String() string {
	current := t.head
	var result string
	for current != nil {
		result += fmt.Sprint(current.Value)
		result += " "
		current = current.Next
	}
	return strings.Trim(result, " ")
}