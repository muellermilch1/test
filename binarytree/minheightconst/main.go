package main

//MinHeightBST constructs a bst from a sorted List
//that has the smallest possible height
//and returns the root of ist
func MinHeightBST(array []int) *BST {
	if len(array) == 0 {
		return nil
	}
	rv := len(array) / 2
	root := &BST{Value: array[rv]}
	//DAS ARRAY IMMER IN ZWEI TEILE SPALTEN
	//die anderen arrays rekursiv aufrufen
	a := array[:rv]
	b := array[rv+1:]
	root.Left = MinHeightBST(a)
	root.Right = MinHeightBST(b)
	return root
}

//BST represents a binary tree
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

//Insert inserts a value in the bst on which it is called
func (tree *BST) Insert(value int) *BST {
	cn := tree
	for {
		if value <= tree.Value {
			if cn.Left != nil {
				cn = cn.Left
				continue
			}
			cn.Left = &BST{Value: value}
			break
		} else {
			if cn.Right != nil {
				cn = cn.Right
				continue
			}
			cn.Right = &BST{Value: value}
			break
		}
	}
	return tree
}

func main() {

}
