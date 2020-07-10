package p4_判断二叉树是否对称

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/omeks/algorithms/datastructure/tree"
)

func Test_isSymmetricTree(t *testing.T) {
	root := &tree.Node{
		Val: 1,
		Left: &tree.Node{
			Val: 2,
			Left: &tree.Node{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.Node{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &tree.Node{
			Val: 2,
			Left: &tree.Node{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	assert.Equal(t, false, isSymmetricTreeIterative(root))
	assert.Equal(t, false, isSymmetricTreeRecursive(root))
	root.Right.Right = &tree.Node{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	assert.Equal(t, true, isSymmetricTreeIterative(root))
	assert.Equal(t, true, isSymmetricTreeRecursive(root))
}
