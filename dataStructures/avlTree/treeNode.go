package avlTree

type treeNode struct {
	Left, Right *treeNode
	Value interface{}
	Height int
}


func (t *treeNode) measureBalance() int {
	var leftHeight int
	var rightHeight int
	if t.Left == nil {
		leftHeight = -1
	} else {
		leftHeight = t.Left.Height 
	}
	if t.Right == nil {
		rightHeight = -1
	} else {
		rightHeight = t.Right.Height 
	}

	return  leftHeight - rightHeight
}