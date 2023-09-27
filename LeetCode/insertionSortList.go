package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	sortedList := &ListNode{
		Val: head.Val,
	}
	node := head.Next
	for node != nil {
		currentNode := sortedList
		var prevNode *ListNode
		for {
			if currentNode.Val > node.Val {
				if prevNode != nil {
					prevNode.Next = &ListNode{
						Val:  node.Val,
						Next: currentNode,
					}
				} else {
					sortedList = &ListNode{
						Val:  node.Val,
						Next: currentNode,
					}
				}
				break
			} else {
				if currentNode.Next == nil {
					currentNode.Next = &ListNode{
						Val: node.Val,
					}
					break
				} else {
					prevNode = currentNode
					currentNode = currentNode.Next
				}
			}
		}
		node = node.Next
	}

	return sortedList
}
