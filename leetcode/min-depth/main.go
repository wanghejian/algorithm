package main

import (
	. "algorithm/tree"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := minDepth(root.Left)
	rd := minDepth(root.Right)

	min := 0
	if ld < rd && ld != 0 {
		min = ld
	} else if rd == 0 {
		min = ld
	} else {
		min = rd
	}

	return min + 1
}
