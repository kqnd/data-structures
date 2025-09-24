package dsa

import "fmt"

type Stack struct {
	Size uint8
	Top *Node
}

// QUICK USAGE
// func main() {
// 	stack := dsa.CreateStack(1)
// 	stack.Add(2)
// 	stack.Add(3)
// 	stack.Add(4)
// 	node := *stack.Pop()
// 	stack.Show()

// 	fmt.Println(node.Value)
// }

func CreateStack(value int) *Stack {
	topNode := Node{
		Value: value,
		Next: nil,
	}
	
	return &Stack{
		Size: 1,
		Top: &topNode,
	}
}

func (s *Stack) Add(value int) {
	newNode := Node{
		Value: value,
		Next: s.Top,
	}

	s.Size++
	s.Top = &newNode
}

func (s *Stack) Pop() *Node {
	if (s.Top == nil) {
		fmt.Println("Stack is empty, can not remove anything")
		return nil
	}

	s.Size--
	removedNode := s.Top
	s.Top = s.Top.Next
	return removedNode
}

func (s *Stack) Show() {
	if s.Top == nil {
		fmt.Println("STACK: empty")
		return 
	}
	
	top := *s.Top
	if top.Next == nil {
		fmt.Printf("STACK[%d]: %d", s.Size, top.Value)
		return
	}

	result := fmt.Sprintf("STACK[%d]: %d -> ", s.Size, top.Value);
	
	for top.Next != nil {
		top = *top.Next
		result += fmt.Sprintf("%d -> ", top.Value)
	}

	result += "None"

	fmt.Println(result)
}
