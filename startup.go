package main

import(
	"fmt"
	ll "./dataStructures/linkedList"
)

func main(){

	linked := ll.ConstructLinkedList(nil)
	linked.AddFirst(12)
	linked.AddFirst(22)

	fmt.Println(linked)
}


