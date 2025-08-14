package solution

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Equals(other *TreeNode) bool {
	if t == nil && other == nil {
		return true
	}
	if t == nil || other == nil {
		return false
	}
	if t.Val != other.Val {
		return false
	}
	return t.Left.Equals(other.Left) && t.Right.Equals(other.Right)
}

func BinaryTreePaths(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
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
