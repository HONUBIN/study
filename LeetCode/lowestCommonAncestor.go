package main

import (
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
   	Val int
   	Left *TreeNode
   	Right *TreeNode
}

// type Stack []*TreeNode

// func (s *Stack) IsEmpty() bool {
// 	return len(*s) == 0
// }

// func (s *Stack) Push(node *TreeNode) {
// 	*s = append(*s, node)
// }

// func (s *Stack) Pop() *TreeNode {
// 	if s.IsEmpty() {
// 		fmt.Println("empty")
// 		return nil
// 	}
// 	top := len(*s)-1
// 	node := (*s)[top]
// 	*s = (*s)[:top]

// 	return node
// }

var count int
var parentsStack = make([]string, 2)
var nodeMap = make(map[int]*TreeNode)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	makeParentsStack("", root, p, q)
	// fmt.Println(parentsStack)

	return findLCA(parentsStack[0], parentsStack[1])
}

func makeParentsStack(stack string, node, p, q *TreeNode) {
	nodeMap[node.Val] = node
	stack = stack + strconv.Itoa(node.Val)
	if node.Val == p.Val || node.Val == q.Val {
		parentsStack[count] = stack
		count++
	}
	if count == 2 {
		return
	} else {
		if node.Left != nil {
			makeParentsStack(stack, node.Left, p, q)
		}
		if node.Right != nil {
			makeParentsStack(stack, node.Right, p, q)
		}
	}
}

func findLCA(pStack, qStack string) *TreeNode {
	leftDepth := len(pStack)
	rightDepth := len(qStack)
	if leftDepth > rightDepth {
		return findLCA(pStack[:rightDepth], qStack)
	} else if leftDepth < rightDepth {
		return findLCA(pStack, qStack[:leftDepth])
	} else if pStack == qStack {
		val, _ := strconv.Atoi(pStack[leftDepth-1:leftDepth])
		return nodeMap[val]
	}
	return findLCA(pStack[:rightDepth-1], qStack[:leftDepth-1])
}
