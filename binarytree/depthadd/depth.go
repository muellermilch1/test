package depth

//BinaryTree is a datasturcture
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

//NodeDepths returns the sum of all depth nodes
func NodeDepths(root *BinaryTree) int {
	return helper(root, 0)
}

func helper(root *BinaryTree, depth int) int {
	sum := depth
	if root.Left != nil {
		sum += helper(root.Left, depth+1)
	}
	if root.Right != nil {
		sum += helper(root.Right, depth+1)
	}
	return sum
}
