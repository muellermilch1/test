package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	dx, px := depthTree(root, nil, x, 0)
	dy, py := depthTree(root, nil, y, 0)
	return (dx == dy) && (px != py)
}

func depthTree(root, parent *TreeNode, x, depth int) (*TreeNode, int) {
	if root.Val == x {
		return parent, depth
	}

	if root.Right != nil {
		p, d := depthTree(root.Right, root, x, depth+1)
		if p != nil {
			return p, d
		}
	}
	if root.Left != nil {
		p, d := depthTree(root.Left, root, x, depth+1)
		if p != nil {
			return p, d
		}
	}
	return nil, -1
}

func main() {
	fmt.Println("fuck")
}
