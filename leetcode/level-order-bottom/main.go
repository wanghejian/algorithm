package main

import (
	"container/list"

	. "algorithm/tree"
)

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var rets [][]int

	l := list.New()
	l.PushBack(root)

	for {
		if l.Len() == 0 {
			break
		}

		var ret []int
		ll := list.New()
		for {
			if l.Len() == 0 {
				break
			}

			e := l.Front()
			t := e.Value.(*TreeNode)
			l.Remove(e)

			if t.Left != nil {
				ll.PushBack(t.Left)
			}
			if t.Right != nil {
				ll.PushBack(t.Right)
			}

			ret = append(ret, t.Val)
		}

		rets = append(rets, ret)
		l = ll
	}

	var rr [][]int
	for i := len(rets) - 1; i >= 0; i-- {
		rr = append(rr, rets[i])
	}

	return rr
}
