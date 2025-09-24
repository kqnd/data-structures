package main

import (
	"fmt"

	"github.com/enjineks/dsa/dsa"
)

func main() {
	stack := dsa.CreateStack(1)
	stack.Add(2)
	stack.Add(3)
	stack.Add(4)
	node := *stack.Pop()
	stack.Show()

	fmt.Println(node.Value)
}