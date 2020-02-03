package main

import (
	"fmt"
	tr "./dataStructures/avlTree"
)

func main(){

	tree := tr.ConstructAVLTree(compare)
	tree.Add(20)
	tree.Add(10)
	tree.Add(15)
	tree.Add(30)
	tree.Add(40)
	tree.Add(7)
	tree.PrintTree()

	fmt.Println("END")
}

func compare(a,b interface{}) int {
	aInt := a.(int)
	bInt := b.(int)
	if aInt > bInt {
		return 1
	} else if aInt < bInt {
		return -1
	}
	return 0
}


