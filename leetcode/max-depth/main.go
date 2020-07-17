package main

import (
	. "algorithm/tree"
	"fmt"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	var max int
	if l > r {
		max = l
	} else {
		max = r
	}

	return max + 1
}
