// t1.Left = t2.Right; t1.Right = t2.Left
// O(n)

package main

import (
	. "algorithm/tree"
	"fmt"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compare(root.Left, root.Right)
}

func compare(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil {
		return false
	} else if t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	if !compare(t1.Left, t2.Right) {
		return false
	}

	if !compare(t1.Right, t2.Left) {
		return false
	}

	return true
}

func main() {
	v1 := []int{1, 2, 3, 0, 0, 4, 0, 0, 2, 4, 0, 0, 3, 0, 0}
	r1, _ := CreateBinaryTree(v1)
	fmt.Println(isSymmetric(r1))

	v2 := []int{1, 2, 0, 3, 0, 0, 2, 0, 3, 0, 0}
	r2, _ := CreateBinaryTree(v2)
	fmt.Println(isSymmetric(r2))
}
