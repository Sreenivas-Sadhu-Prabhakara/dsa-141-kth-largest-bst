package main

import "fmt"

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// KthLargest returns the k-th largest value (1-based) in the BST; ok is
// false when k is out of range. Reverse inorder — right, node, left — walks
// the tree in descending order, stopped after k visits.
func KthLargest(root *TreeNode, k int) (int, bool) {
	if k < 1 {
		return 0, false
	}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right // mirror of the usual spine slide
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return cur.Val, true
		}
		cur = cur.Left
	}
	return 0, false
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(5, tn(3, tn(2, nil, nil), tn(4, nil, nil)), tn(8, tn(6, nil, nil), tn(9, nil, nil)))
	fmt.Println(KthLargest(root, 2))
}
