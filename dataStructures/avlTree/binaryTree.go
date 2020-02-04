package avlTree

import "fmt"

type AVLTree struct { 
	root *treeNode
	compare func(a, b interface{}) int
	count int
	//todo: test Count()o !
}

func ConstructAVLTree(comparator func(a, b interface{}) int) *AVLTree {
	return &AVLTree{ nil, comparator, 0 }
}

func (t *AVLTree) Add(value interface{}) bool {
	newNode := &treeNode {
			Value: value,
			Left: nil,
			Right: nil,
			Height: 0,
		}

	result, ok = t.add(t.root, newNode)
	if ok != nil {
		return false
	} 
	t.root = result
	t.count++
	return true
}

func (t *AVLTree) add(current *treeNode, newNode *treeNode) *treeNode, bool {
	
	if current == nil {
		return newNode
	}

	// result, ok   ??!?!
	if t.compare(current.Value, newNode.Value) > 0 { // go LEFT
		result, ok := t.add(current.Left, newNode)
		if !ok {
			return result, ok
		}
		current.Left = result
	} else if t.compare(current.Value, newNode.Value) < 0 {
		result, ok := t.add(current.Right, newNode)
		if !ok {
			return result, ok
		}
		current.Right = result
	} else {
		return nil, false
	}

	current.Height = setHeight(current)
	balanceFactor := current.measureBalance() 
	var rotatedNode *treeNode
	
	if balanceFactor < -1 {
		// right disbalance
		if current.Right.measureBalance() < 0 { // RR
			rotatedNode = leftRotate(current)
		} else if current.Right.measureBalance() > 0 { // RL
			current.Right = rightRotate(current.Right)
			rotatedNode = leftRotate(current)
		} else {
			//ERROR  
			panic("current.Right.measureBalance() can NOT be equal to 0 when there is Right disbalance measured on current")
		}
	} else if balanceFactor > 1 {
		//left disbalance
		if current.Left.measureBalance() < 0 { // LR
			current.Left = leftRotate(current.Left) 
			rotatedNode = rightRotate(current)
		} else if current.Left.measureBalance() > 0 { // LL
			rotatedNode = rightRotate(current) 
		} else {
			//ERROR 
			panic("current.Left.measureBalance() can NOT be equal to 0 when there is Left disbalance measured on current")
		}
	}

	if rotatedNode != nil {
		return rotatedNode
	}
	//no rotations have been performed
	return current
}

func (t *AVLTree) Remove(val interface{}) (*treeNode, bool) {
	var parent *treeNode
	current := t.root

	// if current is nil ?!?!
	for {
		if current == nil {
			return nil, false
		}

		if t.compare(val, current.Value) > 0 { // Val is bigger (go Right)
			parent = current
			current = current.Right
		} else if t.compare(val, current.Value) < 0 { // Val is smaller (go Left)
			parent = current
			current = current.Left
		} else {
			break
		}
	}

    if current.Left == nil {
    	if current.Right == nil {
    		replaceChild(parent, current, nil)
    	}
    	replaceChild(parent, current, current.Right)
	} else if current.Right == nil { // only Left child    
		replaceChild(parent, current, current.Left)
	} else {
		// case 3 - current has 2 children 

		//find ancestor
		ancestor := current.Right
		var previous *treeNode
		for ancestor.Left != nil {
			ancestor = ancestor.Left
		}


		//ancestorsRight := ancestor.Right

		// here I have the correct ancestor

	}

		//find successor
		// 3 cases -> 
		//successor := current

	

		
	t.count--;
}

//can be improved. Try to move the child's 
//logic inside - if has only right or left ...
func replaceChild(parent, child, replace *treeNode) {
	if parent.Left = child {
		parent.Left = replace
		return 
	}
	parent.Right = replace
}

func setHeight(node *treeNode) int {

	var leftHeight int
	var rightHeight int
	if node.Left == nil {
		leftHeight = -1
	} else {
		leftHeight = node.Left.Height 
	}
	if node.Right == nil {
		rightHeight = -1
	} else {
		rightHeight = node.Right.Height 
	}
	leftHeight++
	rightHeight++
	return max(leftHeight, rightHeight)
}

func leftRotate(node *treeNode) *treeNode {
	child := node.Right
	childsLeft := child.Left
	child.Left = node
	node.Right = childsLeft
	node.Height = setHeight(node)
	child.Height = setHeight(child)
	return child
}

func rightRotate(node *treeNode) *treeNode{
	child := node.Left
	childsRight := child.Right
	child.Right = node
	node.Left = childsRight
	node.Height = setHeight(node)
	child.Height = setHeight(child)
	return child
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func print(root *treeNode){
	if root == nil{
		return
	}
	fmt.Print(root.Value, ", ")
	print(root.Left)
	print(root.Right)
}	

func (t *AVLTree) PrintTree() {
	print(t.root)
}