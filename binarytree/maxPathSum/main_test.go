package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	test := NewBinaryTree(1).insertAll([]int{2, 3, 4, 5, 6, 7})
	actual := MaxPathSum(test)
	if actual != 18 {
		t.Errorf("Failed expected 18 go %v", actual)
	}
}

func TestCase2(t *testing.T) {
	test := &BinaryTree{Value: 1}
	test.Left = &BinaryTree{Value: 2}
	test.Right = &BinaryTree{Value: -1}
	actual := MaxPathSum(test)
	if actual != 3 {
		t.Errorf("Failed expected 3 go %v", actual)
	}
}

func TestCase3(t *testing.T) {
	test := &BinaryTree{Value: 1}
	test.Left = &BinaryTree{Value: -5}
	test.Right = &BinaryTree{Value: 3}
	test.Left.Left = &BinaryTree{Value: 6}
	actual := MaxPathSum(test)
	if actual != 6 {
		t.Errorf("Failed expected 6 go %v", actual)
	}
}

func TestCase4(t *testing.T) {
	test := &BinaryTree{Value: -2}
	actual := MaxPathSum(test)
	if actual != -2 {
		t.Errorf("Failed expected -2 go %v", actual)
	}
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{Value: value}
}

func (tree *BinaryTree) insertAll(values []int) *BinaryTree {
	for _, value := range values {
		tree.insert(value)
	}
	return tree
}

func (tree *BinaryTree) insert(value int) {
	queue := []*BinaryTree{tree}
	var current *BinaryTree
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = NewBinaryTree(value)
			break
		}
		queue = append(queue, current.Left)
		if current.Right == nil {
			current.Right = NewBinaryTree(value)
			break
		}
		queue = append(queue, current.Right)
	}
}
