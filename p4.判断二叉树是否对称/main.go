package p4_判断二叉树是否对称

import (
	"github.com/omeks/algorithms/datastructure"
)

func isSymmetric(left *datastructure.TreeNode, right *datastructure.TreeNode) bool {
	if left != nil && right != nil {
		return isSymmetric(left.Left, right.Right) && isSymmetric(left.Right, right.Left)
	}
	return left == nil && right == nil
}

// Time: O(n), Space: O(n)
func isSymmetricTreeRecursive(root *datastructure.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}

// Time: O(n), Space: O(n)
func isSymmetricTreeIterative(root *datastructure.TreeNode) bool {
	if root == nil {
		return true
	}
	stack := datastructure.NewStack()
	stack.Push(root.Left)
	stack.Push(root.Right)
	for !stack.IsEmpty() {
		n1 := stack.Pop()
		n2 := stack.Pop()

		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}

		left := n1.(*datastructure.TreeNode)
		right := n2.(*datastructure.TreeNode)

		if left.Val != right.Val {
			return false
		}

		if left.Left != nil {
			stack.Push(left.Left)
		}
		if right.Right != nil {
			stack.Push(right.Right)
		}
		if left.Right != nil {
			stack.Push(left.Right)
		}
		if right.Left != nil {
			stack.Push(right.Left)
		}
	}
	return true
}
