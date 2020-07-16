package datastructure

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr ...int) *ListNode {
	if len(arr) == 0 {
		return &ListNode{Val: 0, Next: nil}
	}

	head := &ListNode{Val: arr[0], Next: nil}
	for i := 1; i < len(arr); i++ {
		head.Insert(arr[i])
	}

	return head
}

func (p *ListNode) Insert(val int) {
	if p.Next != nil {
		p.Next.Insert(val)
	} else {
		p.Next = &ListNode{Val: val, Next: nil}
	}
}

// List2Array - convert list to array
func List2Array(head *ListNode) (arr []int) {
	if head == nil {
		return nil
	}

	p := head
	for p != nil {
		arr = append(arr, p.Val)
		p = p.Next
	}

	return
}

// Array2List - convert array to list
func Array2List(arr []int) (head *ListNode, tail *ListNode) {
	if len(arr) == 0 {
		return nil, nil
	}

	head = &ListNode{Val: arr[0], Next: nil}
	p := head
	for i := 1; i < len(arr); i++ {
		p, p.Next = p.Next, &ListNode{Val: arr[i], Next: nil}
	}
	tail = p

	return
}
