package dsa

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

type Tree struct {
	Size int
	Root *TreeNode
}

func CreateTree(value int) *Tree {
	return &Tree{
		Size: 1, 
		Root: &TreeNode{
			Value: value,
			Left: nil,
			Right: nil,
		},
	}
}

func (t *Tree) Invert() {
	invertRecursively(t.Root)
}

func invertRecursively(node *TreeNode) {
	if node == nil {
		return
	}

	temp := node.Left
	node.Left = node.Right
	node.Right = temp

	invertRecursively(node.Left)
	invertRecursively(node.Right)
}


func (t *Tree) Add(value int) {
	addRecursively(t.Root, value)
	t.Size++
}

func addRecursively(node *TreeNode, value int) {
	if node == nil {
		return
	}

	newNode := &TreeNode{
		Value: value,
		Left: nil,
		Right: nil,
	}

	nodeValue := node.Value
	if value > nodeValue {
		if node.Right == nil {
			node.Right = newNode
			return
		} else {
			addRecursively(node.Right, value)
		}
	} 
	
	if value < nodeValue {
		if node.Left == nil {
			node.Left = newNode
			return
		} else {
			addRecursively(node.Left, value)
		}
	}

}

func printNode(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	printNode(node.Right, depth+1)

	fmt.Printf("%s%d\n", strings.Repeat("    ", depth), node.Value)

	printNode(node.Left, depth+1)
}


func (t *Tree) Show() {
	if t.Root == nil {
		fmt.Println("Tree is empty")
		return
	}
	printNode(t.Root, 0)
}
