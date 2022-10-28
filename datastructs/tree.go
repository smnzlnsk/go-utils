package datastructs

import (
	"fmt"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func NewNode(value int) *Tree {
	t := new(Tree)
	t.Value = value
	return t
}

func PrintTree(t *Tree) {
	if l := t.Left; l != nil {
		PrintTree(l)
	}
	fmt.Print(t.Value, " ")
	if r := t.Right; r != nil {
		PrintTree(r)
	}
}

func InsertBinTree(t *Tree, n *Tree) *Tree {
	if t == nil {
		return n
	}
	if t.Value >= n.Value {
		if t.Left != nil {
			InsertBinTree(t.Left, n)
		} else {
			t.Left = n
		}
	}
	if t.Value < n.Value {
		if t.Right != nil {
			InsertBinTree(t.Right, n)
		} else {
			t.Right = n
		}
	}
	return t
}
// TODO check for parent node to fulfil binary-tree condition
func CorrectBinTree(t *Tree) bool {
	var leftCorrect, rightCorrect bool
	if l := t.Left; l != nil {
		if l.Value <= t.Value {
			fmt.Println("left:", l.Value, t.Value)
			leftCorrect = CorrectBinTree(l)
		} else {
			return false
		}
	} else {
		leftCorrect = true
	}
	if r := t.Right; r != nil {
		if r.Value > t.Value {
			fmt.Println("right:", r.Value, t.Value)
			rightCorrect = CorrectBinTree(r)
		} else {
			return false
		}
	} else {
		rightCorrect = true
	}
	return leftCorrect && rightCorrect
}
