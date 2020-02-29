package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// o(N)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h1, h2 := l1, l2

	var (
		root, pre *ListNode
		carry     int
	)

	for {
		curr := &ListNode{}
		if h1 != nil && h2 != nil {
			curr.Val = h1.Val + h2.Val + carry

			if curr.Val >= 10 {
				curr.Val %= 10
				carry = 1
			} else {
				carry = 0
			}

			h1 = h1.Next
			h2 = h2.Next
		} else if h1 != nil {
			curr.Val = h1.Val + carry
			if curr.Val >= 10 {
				curr.Val %= 10
				carry = 1
			} else {
				carry = 0
			}
			h1 = h1.Next
		} else if h2 != nil {
			curr.Val = h2.Val + carry
			if curr.Val >= 10 {
				curr.Val %= 10
				carry = 1
			} else {
				carry = 0
			}
			h2 = h2.Next
		} else {
			if carry != 0 {
				curr.Val = carry
				carry = 0
			} else {
				break
			}
		}

		if root == nil {
			root = curr
		}

		if pre != nil {
			pre.Next = curr
		}
		pre = curr

	}

	return root
}

func newList(num int) *ListNode {
	var root *ListNode
	var pre *ListNode
	for {
		tmp := num % 10
		num = num / 10

		curr := &ListNode{}
		curr.Val = tmp

		if root == nil {
			root = curr
		}

		if pre != nil {
			pre.Next = curr
		}
		pre = curr

		if num == 0 {
			break
		}
	}

	return root
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
	fmt.Println()
}

func main() {
	l1 := newList(342)
	l2 := newList(465)
	printList(addTwoNumbers(l1, l2))

	l1 = newList(0)
	l2 = newList(989)
	printList(addTwoNumbers(l1, l2))

	l1 = newList(5)
	l2 = newList(5)
	printList(addTwoNumbers(l1, l2))

	l1 = newList(1)
	l2 = newList(99)
	printList(addTwoNumbers(l1, l2))

	l1 = newList(1234)
	l2 = newList(456789)
	printList(addTwoNumbers(l1, l2))
}
