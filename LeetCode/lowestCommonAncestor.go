package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []*TreeNode

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() {
		fmt.Println("empty")
		return nil
	}
	top := len(*s) - 1
	node := (*s)[top]
	*s = (*s)[:top]

	return node
}

var count int
var parentsStack []Stack
var stack Stack

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	count = 0
	parentsStack = make([]Stack, 2)
	stack = Stack{}
	makeParentsStack(root, p, q, stack)

	return findLCA(parentsStack[0], parentsStack[1])
}

func makeParentsStack(node, p, q *TreeNode, stack Stack) {
	stack.Push(node)
	if node.Val == p.Val || node.Val == q.Val {
		copiedStack := make([]*TreeNode, len(stack))
		copy(copiedStack, stack)
		parentsStack[count] = copiedStack
		count++
	}
	if count == 2 {
		return
	} else {
		if node.Left != nil {
			makeParentsStack(node.Left, p, q, stack)
		}
		if node.Right != nil {
			makeParentsStack(node.Right, p, q, stack)
		}
	}
}

func findLCA(pStack, qStack Stack) *TreeNode {
	pDepth := len(pStack)
	qDepth := len(qStack)
	if pDepth > qDepth {
		pStack.Pop()
		return findLCA(pStack, qStack)
	} else if pDepth < qDepth {
		qStack.Pop()
		return findLCA(pStack, qStack)
	} else {
		p := pStack.Pop()
		q := qStack.Pop()
		if p == q {
			return p
		}
		return findLCA(pStack, qStack)
	}
}
