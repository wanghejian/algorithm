package tree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateBinaryTree(vals []int) (*TreeNode, []int) {
	if len(vals) == 0 {
		return nil, nil
	}

	if vals[0] == 0 {
		return nil, vals[1:]
	}

	t := &TreeNode{}
	t.Val = vals[0]

	l, v1 := CreateBinaryTree(vals[1:])
	r, v2 := CreateBinaryTree(v1)

	t.Left = l
	t.Right = r

	return t, v2
}

func PreOrderTraverse(t *TreeNode) {
	if t == nil {
		return
	}

	fmt.Println(t.Val)
	PreOrderTraverse(t.Left)
	PreOrderTraverse(t.Right)
}
