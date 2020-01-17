package main

import(
	"fmt"
	ll "./dataStructures/linkedList"
)

func main(){

	linked := ll.ConstructLinkedList()
	linked.AddFirst(12)

	ex := ll.Exported{14}
	fmt.Println(ex)
}


