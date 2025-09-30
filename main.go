package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	lista1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	appendNode(lista1, 2)
	appendNode(lista1, 4)

	lista2 := &ListNode{
		Val:  1,
		Next: nil,
	}

	appendNode(lista2, 3)
	appendNode(lista2, 4)

	result := mergeTwoLists(lista1, lista2)
	current := result
	for current.Next != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
	fmt.Println(current.Val)
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

func appendNode(node *ListNode, val int) {
	current := node
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Val: val}
}
