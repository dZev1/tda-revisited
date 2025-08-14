package main

import (
	"backtracking/solution"
	"fmt"
)

func main() {
	tree := solution.TreeNode{
		Val: 1,
		Left: &solution.TreeNode{
			Val: 2,
			Right: &solution.TreeNode{
				Val: 5,
			},
		},
		Right: &solution.TreeNode{
			Val: 3,
		},
	}

	fmt.Println(solution.BinaryTreePaths(&tree))
}