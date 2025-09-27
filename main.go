package main

import (
	"github.com/kqnd/dsa/dsa"
)

func main() {
	tree := dsa.CreateTree(5)
	tree.Add(1)
	tree.Add(3)
	tree.Show()
}
