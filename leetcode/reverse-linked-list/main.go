package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var (
		pre, cur, post *ListNode
	)

	cur = head
	for cur != nil {
		post = cur.Next
		cur.Next = pre
		pre = cur
		cur = post
	}

	return pre
}

func newList(nums ...int) *ListNode {
	var (
		root, pre *ListNode
	)

	for _, num := range nums {
		node := &ListNode{}
		node.Val = num

		if root == nil {
			root = node
		}

		if pre != nil {
			pre.Next = node
		}
		pre = node
	}

	return root
}

func printList(l *ListNode) {
	n := l
	for n != nil {
		fmt.Print(n.Val, " ")
		n = n.Next
	}

	fmt.Println()
}

func main() {
	l := newList(4, 5, 6, 1, 2, 3)
	printList(l)
	l = reverseList(l)
	printList(l)

	l = newList(1)
	l = reverseList(l)
	printList(l)
}
