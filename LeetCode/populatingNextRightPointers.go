package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	levels := make([][]*Node, 6000)
	level := 0
	levels[level] = append(levels[level], root)
	levels[level+1] = append(levels[level+1], root.Left)
	levels[level+1] = append(levels[level+1], root.Right)

	for true {
		level += 1
		notNull := false
		var prev *Node
		for _, n := range levels[level] {
			if n == nil {
				continue
			}
			if prev != nil {
				prev.Next = n
			}
			if n.Left != nil || n.Right != nil {
				notNull = true
			}
			levels[level+1] = append(levels[level+1], n.Left)
			levels[level+1] = append(levels[level+1], n.Right)
			prev = n
		}

		if !notNull {
			break
		}
	}

	return root
}
