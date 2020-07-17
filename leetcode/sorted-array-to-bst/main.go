package main

import (
	. "algorithm/tree"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	pivot := len(nums) / 2

	var l, r *TreeNode

	if pivot != 0 {
		l = sortedArrayToBST(nums[:pivot])
	}

	if pivot+1 < len(nums) {
		r = sortedArrayToBST(nums[pivot+1:])
	}

	t := &TreeNode{Val: nums[pivot]}
	t.Left = l
	t.Right = r
	return t
}
