package main

import "github.com/enjineks/dsa/dsa"

func main() {
	tree := dsa.CreateTree(5)
	tree.Add(1)
	tree.Add(3)
	tree.Show()
}