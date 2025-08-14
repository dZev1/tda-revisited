package solution

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreePaths(root *TreeNode) []string {
	paths := []string{}

	var backtrack func(node *TreeNode, path string)
	backtrack = func(node *TreeNode, path string) {
		leaf := node.Left == nil && node.Right == nil
		if leaf {
			paths = append(paths, fmt.Sprintf("%v%v", path, node.Val))
		}
		currPath := fmt.Sprintf("%v%v->", path, node.Val)
		if node.Left != nil {
			backtrack(node.Left, currPath)
		}
		if node.Right != nil {
			backtrack(node.Right, currPath)
		}
	}
	backtrack(root, "")
	return paths
}
