// 通过遍历在比较相等性，可以判断两棵二叉树是不是一样
// O(n)

package main

import (
	. "algorithm/tree"
	"fmt"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return inOrderTraverse(p, q)
}

func inOrderTraverse(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil {
		return false
	} else if q == nil {
		return false
	}

	if !inOrderTraverse(p.Left, q.Left) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !inOrderTraverse(p.Right, q.Right) {
		return false
	}

	return true
}

func main() {
	// pre order list
	v1 := []int{1, 2, 0, 3, 0, 0, 4, 0, 0}
	r1, _ := CreateBinaryTree(v1)
	fmt.Println(isSameTree(r1, r1))

	v2 := []int{1, 2, 3}
	r2, _ := CreateBinaryTree(v2)
	fmt.Println(isSameTree(r2, r2))

	v3 := []int{1, 2}
	r3, _ := CreateBinaryTree(v3)
	v4 := []int{1, 0, 2}
	r4, _ := CreateBinaryTree(v4)
	fmt.Println(isSameTree(r3, r4))

	v5 := []int{1, 2, 1}
	r5, _ := CreateBinaryTree(v5)
	v6 := []int{1, 1, 2}
	r6, _ := CreateBinaryTree(v6)
	fmt.Println(isSameTree(r5, r6))
}
