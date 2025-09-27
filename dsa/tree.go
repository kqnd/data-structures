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

func (t *Tree) Remove(value int) {
	t.Root = removeRecursively(t.Root, value)
	t.Size--
}

func removeRecursively(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value > node.Value {
		node.Right = removeRecursively(node.Right, value)
	} else if value < node.Value {
		node.Left = removeRecursively(node.Left, value)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		}
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		successor := findMin(node.Right)
		node.Value = successor.Value
		node.Right = removeRecursively(node.Right, successor.Value)
	}
	return node
}

func (t *Tree) Search(value int) *TreeNode {
	return searchRecursively(t.Root, value)
}

func searchRecursively(node *TreeNode, value int) *TreeNode {
	if node == nil || node.Value == value {
		return node
	}

	if value > node.Value {
		return searchRecursively(node.Right, value)
	} else {
		return searchRecursively(node.Left, value)
	}
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func addRecursively(node *TreeNode, value int) {

	newNode := &TreeNode{
		Value: value,
		Left: nil,
		Right: nil,
	}

	nodeValue := node.Value

	if value == nodeValue {
		temp := node.Left
		node.Left = newNode
		newNode.Left = temp
		return
	}

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
