package dsa

import "fmt"

type Queue struct {
	Size uint8
	Tail *Node
}


// QUICK USAGE
// func main() {
// 	queue := dsa.CreateQueue(1)
// 	queue.Add(2)
// 	queue.Add(3)
// 	queue.Remove()
// 	queue.Show()
// }

func CreateQueue(value int) *Queue {
	return &Queue{
		Size: 1,
		Tail: &Node{
			Value: value,
			Next: nil,
		},
	}
}

func (q *Queue) Add(value int) {
	newNode := &Node{
		Value: value,
		Next: nil,
	}
	
	currentNode := q.Tail
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	q.Size++
	currentNode.Next = newNode
}

func (q *Queue) Remove() *Node {
	if q.Tail == nil {
		fmt.Println("Queue is empty")
		return nil
	}

	q.Size--
	removedNode := q.Tail
	q.Tail = q.Tail.Next
	return removedNode
}

func (q *Queue) Show() {
	if q.Tail == nil {
		fmt.Println("QUEUE: empty")
		return 
	}
	
	tail := *q.Tail
	if tail.Next == nil {
		fmt.Printf("QUEUE[%d]: %d", q.Size, tail.Value)
		return
	}

	result := fmt.Sprintf("QUEUE[%d]: %d -> ", q.Size, tail.Value);
	
	for tail.Next != nil {
		tail = *tail.Next
		result += fmt.Sprintf("%d -> ", tail.Value)
	}

	result += "None"

	fmt.Println(result)
}
