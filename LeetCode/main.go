package main

import "fmt"

func main() {
	fmt.Println(lowestCommonAncestor(
		&TreeNode{
			Val: -1,
			Left: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -2,
					Left: &TreeNode{
						Val: 8,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}, &TreeNode{
			Val: 3,
		}, &TreeNode{
			Val: 8,
		}))
}
