package linkedList

import "fmt"

type listNode struct{
	Next *listNode 
	Previous *listNode 
	Value interface{}
}

func (target listNode) String() string {
	return fmt.Sprintf("%v", target.Value)
}