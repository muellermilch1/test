package main

import "math"

//BinaryTree represents a binary tree
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

//MaxPathSum returns max path sum
func MaxPathSum(tree *BinaryTree) int {
	_, v := helper(tree, math.MinInt64)
	return v
}

func helper(tree *BinaryTree, max int) (int, int) {
	var a, b, m int
	if tree.Right != nil {
		a, m = helper(tree.Right, max)
		max = bigger(m, max)
	}
	if tree.Left != nil {
		b, m = helper(tree.Left, max)
		max = bigger(m, max)
	}
	mpath := bigger(a+tree.Value, b+tree.Value)
	max = bigger(tree.Value, tree.Value+a+b, mpath, max)

	return mpath, max
}

func bigger(args ...int) int {
	max := math.MinInt64
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {

}
