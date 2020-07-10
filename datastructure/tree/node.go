package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Constructor(val int) *Node {
	return &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (p *Node) Insert(val int) {
	if val < p.Val {
		if p.Left != nil {
			p.Left.Insert(val)
		} else {
			p.Left = &Node{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		}
	} else {
		if p.Right != nil {
			p.Right.Insert(val)
		} else {
			p.Right = &Node{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		}
	}
}

func (p *Node) PreOrderRecursive() []int {
	var res []int
	res = append(res, p.Val)
	if p.Left != nil {
		res = append(res, p.Left.PreOrderRecursive()...)
	}
	if p.Right != nil {
		res = append(res, p.Right.PreOrderRecursive()...)
	}
	return res
}

func (p *Node) InOrderRecursive() []int {
	var res []int
	if p.Left != nil {
		res = append(res, p.Left.InOrderRecursive()...)
	}
	res = append(res, p.Val)
	if p.Right != nil {
		res = append(res, p.Right.InOrderRecursive()...)
	}
	return res
}

func (p *Node) PostOrderRecursive() []int {
	var res []int
	if p.Left != nil {
		res = append(res, p.Left.PostOrderRecursive()...)
	}
	if p.Right != nil {
		res = append(res, p.Right.PostOrderRecursive()...)
	}
	res = append(res, p.Val)
	return res
}

func (p *Node) LevelOrderIterative() []int {
	var res []int
	queue := []*Node{p}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			res = append(res, node.Val)
		}
		queue = queue[size:]
	}
	return res
}

func (p *Node) PreOrderIterative() []int {
	var (
		res   []int
		stack []*Node
	)
	node := p
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			res = append(res, node.Val)
			node = node.Left
		}
		if len(stack) > 0 {
			last := len(stack) - 1
			node = stack[last].Right
			stack = stack[:last]
		}
	}
	return res
}

func (p *Node) InOrderIterative() []int {
	var (
		res   []int
		stack []*Node
	)
	node := p
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) > 0 {
			last := len(stack) - 1
			res = append(res, stack[last].Val)
			node = stack[last].Right
			stack = stack[:last]
		}
	}
	return res
}

func (p *Node) PostOrderIterative() []int {
	var (
		res   []int
		stack []*Node
	)
	node := p
	lastVisit := p
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		last := len(stack) - 1
		// 查看当前栈顶元素
		node = stack[last]
		// 如果其右子树也为空，或者右子树已经访问，则可以直接输出当前节点的值
		if node.Right == nil || node.Right == lastVisit {
			res = append(res, node.Val)
			stack = stack[:last]
			lastVisit = node
			node = nil
		} else {
			// 否则继续遍历右子树
			node = node.Right
		}
	}
	return res
}

func (p *Node) PostOrderIterative2() []int {
	var (
		res   []int
		stack []*Node
	)
	node := p
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			res = append(res, node.Val)
			node = node.Right
		}
		if len(stack) > 0 {
			last := len(stack) - 1
			node = stack[last].Left
			stack = stack[:last]
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
