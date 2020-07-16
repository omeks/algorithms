package p4_判断二叉树是否对称

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/omeks/algorithms/datastructure"
)

func Test_isSymmetricTree(t *testing.T) {
	root := &datastructure.TreeNode{
		Val: 1,
		Left: &datastructure.TreeNode{
			Val: 2,
			Left: &datastructure.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &datastructure.TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &datastructure.TreeNode{
			Val: 2,
			Left: &datastructure.TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	assert.Equal(t, false, isSymmetricTreeIterative(root))
	assert.Equal(t, false, isSymmetricTreeRecursive(root))
	root.Right.Right = &datastructure.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	assert.Equal(t, true, isSymmetricTreeIterative(root))
	assert.Equal(t, true, isSymmetricTreeRecursive(root))
}
