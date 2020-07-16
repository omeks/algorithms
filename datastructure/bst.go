package datastructure

type BinarySearchTree struct {
	Root *TreeNode
}

func NewBinarySearchTree(arr ...int) *BinarySearchTree {
	return &BinarySearchTree{Root: NewTreeNode(arr...)}
}

func (p *BinarySearchTree) Insert(val int) {
	p.Root.Insert(val)
}

func (p *BinarySearchTree) InOrder() []int {
	return p.Root.InOrderRecursive()
}

func (p *BinarySearchTree) Contains(val int) bool {
	node := p.Root
	for {
		if node == nil {
			return false
		} else if node.Val == val {
			return true
		} else if node.Val > val {
			node = node.Left
		} else if node.Val < val {
			node = node.Right
		}
	}
}

func (p *BinarySearchTree) Min() int {
	node := p.Root
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}

func (p *BinarySearchTree) Max() int {
	node := p.Root
	for node.Right != nil {
		node = node.Right
	}
	return node.Val
}
