package main

import (
	. "algorithm/tree"
)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return hps(root, sum)
}
func hps(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val
	if sum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if hps(root.Left, sum) || hps(root.Right, sum) {
		return true
	}

	return false
}
