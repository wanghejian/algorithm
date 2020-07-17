package main

import (
	. "algorithm/tree"
)

// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	if _, ok := maxDepth(root); ok {
// 		return true
// 	}

// 	return false
// }

// func maxDepth(t *TreeNode) (int, bool) {
// 	if t == nil {
// 		return 0, true
// 	}

// 	ld, ok1 := maxDepth(t.Left)
// 	rd, ok2 := maxDepth(t.Right)

// 	if ld-rd > 1 || rd-ld > 1 {
// 		return 0, false
// 	}

// 	if !ok1 || !ok2 {
// 		return 0, false
// 	}

// 	max := 0
// 	if ld > rd {
// 		max = ld
// 	} else {
// 		max = rd
// 	}

// 	return max + 1, true
// }

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)

	if ld-rd > 1 || rd-ld > 1 {
		return false
	}

	if !isBalanced(root.Left) {
		return false
	}

	if !isBalanced(root.Right) {
		return false
	}

	return true
}

func maxDepth(t *TreeNode) int {
	if t == nil {
		return 0
	}

	ld := maxDepth(t.Left)
	rd := maxDepth(t.Right)

	max := 0
	if ld > rd {
		max = ld
	} else {
		max = rd
	}

	return max + 1
}
