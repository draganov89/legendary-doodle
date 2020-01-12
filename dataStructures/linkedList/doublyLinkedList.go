package linkedList


//LinkedList implements a Doubly Linked List data structure
type LinkedList struct{
	head *listNode
	tail *listNode
	count int
}

// Count gets the number of elements contained in the Linked List
func (target LinkedList) Count() int{
	return target.count
}

// AddFirst inserts element at the beginning of the Linked List
func (target *LinkedList) AddFirst(element interface{}) {
	newNode := &listNode{nil, nil, element}

	if (target.head == nil){ // linkedList is Empty
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

	if (target.tail == nil){
		target.AddFirst(element)
		return
	}

	target.tail.Next = newNode
	newNode.Previous = target.tail
	target.tail = newNode
	target.count++
}

//String method returns a string representing the Linked List
func (target LinkedList) String() string {
	temp := target.head
	var result string
	for (temp != nil) {
		result += temp.String()
		result += " "
		temp = temp.Next
	}
	return result
}