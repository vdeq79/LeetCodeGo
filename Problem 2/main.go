package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(a int, b int) (int, int) {
	sum := a + b
	return sum % 10, sum / 10
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carrier int = 0
	var Head *ListNode = l1
	l2Longer := false

	for l1.Next != nil {
		if l2 != nil {
			l1.Val, carrier = add(l1.Val, l2.Val+carrier)
			l2 = l2.Next
		} else {
			l1.Val, carrier = add(l1.Val, carrier)
		}

		l1 = l1.Next
	}

	for l2 != nil {
		l1.Val, carrier = add(l1.Val, l2.Val+carrier)
		l2 = l2.Next
		l2Longer = true

		if l2 != nil {
			l1.Next = &ListNode{Val: 0}
			l1 = l1.Next
		}
	}

	if carrier == 1 {
		if l2Longer {
			l1.Next = &ListNode{Val: 1}
		} else {
			l1.Val, carrier = add(l1.Val, carrier)

			if carrier == 1 {
				l1.Next = &ListNode{Val: 1}
			}
		}
	}

	return Head
}

func main() {
	var result = addTwoNumbers(&ListNode{Val: 5}, &ListNode{Val: 5})

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
