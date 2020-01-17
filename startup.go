package main

import(
	"fmt"
	ll "./dataStructures/linkedList"
)

func main(){

	linked := ll.ConstructLinkedList(nil)
	linked.AddFirst(12)

	fmt.Println(linked)
}


