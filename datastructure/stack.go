package datastructure

import (
	"sync"
)

type (
	Stack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func NewStack(arr ...int) *Stack {
	if len(arr) == 0 {
		return &Stack{nil, 0, &sync.RWMutex{}}
	}

	stack := &Stack{&node{value: arr[0]}, 1, &sync.RWMutex{}}
	for i := 1; i < len(arr); i++ {
		stack.Push(arr[i])
	}

	return stack
}

func (this *Stack) Len() int {
	return this.length
}

func (this *Stack) IsEmpty() bool {
	return this.length == 0
}

func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

func (this *Stack) Push(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func (this *Stack) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
