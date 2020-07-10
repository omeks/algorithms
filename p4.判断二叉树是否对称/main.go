package p4_判断二叉树是否对称

import (
	"container/list"

	"github.com/omeks/algorithms/datastructure/tree"
)

func isSymmetric(left *tree.Node, right *tree.Node) bool {
	if left != nil && right != nil {
		return isSymmetric(left.Left, right.Right) && isSymmetric(left.Right, right.Left)
	}
	return left == nil && right == nil
}

// Time: O(n), Space: O(n)
func isSymmetricTreeRecursive(root *tree.Node) bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}

// Time: O(n), Space: O(n)
func isSymmetricTreeIterative(root *tree.Node) bool {
	list.New()
	if root == nil {
		return true
	}
	var (
		stack []*tree.Node
	)
	stack = append(stack, root.Left, root.Right)
	for len(stack) > 0 {
		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		right := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}

		stack = append(stack, left.Left)
		stack = append(stack, right.Right)
		stack = append(stack, left.Right)
		stack = append(stack, right.Left)
	}
	return true
}
