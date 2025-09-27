package dsa

import (
	"fmt"
	"strings"

	"github.com/kqnd/dsa/utils"
)

type Book struct {
	ID int
	Title string
}

type LeafTreeNode struct {
	Key int
	Left, Right *LeafTreeNode
	Leaf *Book
}

type LeafTree struct {
	Root *LeafTreeNode
}

var rootKey = utils.Generate(200, 1000)

func CreateLeafTree() *LeafTree {
	return &LeafTree{
		Root: &LeafTreeNode{
			Key: rootKey,
			Left: nil,
			Right: nil,
			Leaf: nil,
		},
	}
}

func CreateBook(title string) *Book {
	bookID := utils.Generate(1000, 9999)
	fmt.Printf("Creating book: %s [#%d]\n", title, bookID)
	return &Book{
		ID: bookID,
		Title: title,
	}
}

func (t *LeafTree) Add(book *Book) *LeafTreeNode {
	key := utils.Generate(1, 2000)
	node := addLeafRecursively(t.Root, key, book)
	t.Root = node
	return t.Find(key)
}

func addLeafRecursively(node *LeafTreeNode, key int, book *Book) *LeafTreeNode {
	if node == nil {
		newNode := LeafTreeNode{
			Key: key,
			Leaf: book,
		}
		if book != nil {
			fmt.Printf("Added book: %s [#%d] with key %d\n", book.Title, book.ID, key)
		}
		return &newNode
	}

	if key < node.Key {
		node.Left = addLeafRecursively(node.Left, key, book)
	} else {
		node.Right = addLeafRecursively(node.Right, key, book)
	}

	return node
}

func (lt *LeafTree) Find(key int) *LeafTreeNode {
	return findRecursively(lt.Root, key)
}

func (lt *LeafTree) FindByBookID(id int) *LeafTreeNode {
	return findByBookID(lt.Root, id)
}

func findByBookID(node *LeafTreeNode, id int) *LeafTreeNode {
	if node == nil {
		return node
	}

	if node.Leaf != nil && node.Leaf.ID == id {
		return node
	}

	if left := findByBookID(node.Left, id); left != nil {
		return left
	}

	return findByBookID(node.Right, id)
}

func findRecursively(node *LeafTreeNode, key int) *LeafTreeNode {
	if node == nil || node.Key == key {
		return node
	}
	
	if key < node.Key {
		return findRecursively(node.Left, key)
	} else {
		return findRecursively(node.Right, key)
	}
}

func printLeafNode(node *LeafTreeNode, depth int) {
	if node == nil {
		return
	}

	leafInfo := ""
	if node.Leaf != nil {
		leafInfo = fmt.Sprintf(" [%s #%d]", node.Leaf.Title, node.Leaf.ID)
	}

	printLeafNode(node.Right, depth+1)

	fmt.Printf("%s%d%s\n", strings.Repeat("    ", depth), node.Key, leafInfo)

	printLeafNode(node.Left, depth+1)
}


func (lt *LeafTree) Show() {
	if lt.Root == nil {
		fmt.Println("Tree is empty")
		return
	}
	printLeafNode(lt.Root, 0)
}
