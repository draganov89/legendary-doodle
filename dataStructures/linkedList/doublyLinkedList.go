package linkedList

import(
	"strings"
	"errors"
)

//LinkedList implements a Doubly Linked List data structure
type LinkedList struct{
	head, tail *listNode
	compare func(a, b interface{}) int
	count int
}

// Count gets the number of elements contained in the Linked List
func (target LinkedList) Count() int{
	return target.count
}
//ConstructLinkedList is a constructor function for Linked List
func ConstructLinkedList(comparator func(a, b interface{}) int) *LinkedList{
	return &LinkedList{nil, nil, comparator, 0}
}

// AddFirst inserts element at the beginning of the Linked List
func (target *LinkedList) AddFirst(element interface{}) {
	newNode := &listNode{nil, nil, element}

	// linkedList is Empty
	if target.head == nil{ 
		target.head = newNode
		target.tail = newNode
	} else {
		newNode.Next = target.head
		target.head.Previous = newNode
		target.head = newNode
	}
	target.count++
}

// AddLast inserts element at the end of the Linked List
func (target *LinkedList) AddLast(element interface{}) {
	newNode := &listNode{nil, nil, element}

	if target.tail == nil {
		target.AddFirst(element)
		return
	}

	target.tail.Next = newNode
	newNode.Previous = target.tail
	target.tail = newNode
	target.count++
}

// RemoveFirst method removes the first element of the LinkedList and returns it. If the LinkedList is empty error will be returned.
func (target *LinkedList) RemoveFirst() (interface{}, error){
	if target.head == nil {
		return nil, errors.New("target Linked List is empty!")
	}

	removed := target.head
	target.head = target.head.Next

	if target.head == nil{
		target.tail = nil
	} else{
		target.head.Previous = nil
	}

	target.count--
	return removed.Value, nil
}

// RemoveLast method removes the last element of the LinkedList and returns it. If the LinkedList is empty error will be returned.
func (target *LinkedList) RemoveLast()(interface{}, error){
	if target.tail == nil {
		return nil, errors.New("target Linked List is empty!") 
	}

	removed := target.tail
	target.tail = target.tail.Previous

	if target.tail == nil {
		target.head = nil
	} else {
		target.tail.Next = nil
	}

	target.count--
	return removed.Value, nil
}

// func (target LinkedList) Contains(value Comparable) bool {
// 	temp := target.head
// 	for temp != nil {
// 		if temp.Value.CompareTo(value) == 0 {
// 			return true
// 		}
// 	}
// 	return false
// }


//String method returns a string representing the Linked List
func (target LinkedList) String() string {
	temp := target.head
	var result string
	for (temp != nil) {
		result += temp.String()
		result += " "
		temp = temp.Next
	}
	
	return strings.Trim(result, " ")
}

